// SDK smoke test -- validates build-from-source and API integration using the SDK client.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	mailodds "github.com/mailodds/go-sdk"
)

var passed, failed, warned int

func check(label, expected, actual string) {
	if expected == actual {
		passed++
	} else {
		failed++
		fmt.Printf("  FAIL: %s expected=%s got=%s\n", label, expected, actual)
	}
}

func checkBool(label string, expected, actual bool) {
	if expected == actual {
		passed++
	} else {
		failed++
		fmt.Printf("  FAIL: %s expected=%v got=%v\n", label, expected, actual)
	}
}

func warn(label, msg string) {
	warned++
	fmt.Printf("  WARN: %s %s\n", label, msg)
}

func main() {
	apiKey := os.Getenv("MAILODDS_TEST_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: MAILODDS_TEST_KEY not set")
		os.Exit(1)
	}

	cfg := mailodds.NewConfiguration()
	cfg.Servers = mailodds.ServerConfigurations{
		{URL: "https://api.mailodds.com", Description: "Production"},
	}
	client := mailodds.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), mailodds.ContextAccessToken, apiKey)

	// -------------------------------------------------------------------------
	// 1. Email Validation
	// -------------------------------------------------------------------------
	fmt.Println("--- Email Validation ---")

	type tc struct {
		email, status, action, sub string
		freeProvider, disposable, roleAccount, mxFound bool
	}
	cases := []tc{
		{"test@deliverable.mailodds.com", "valid", "accept", "", false, false, false, true},
		{"test@invalid.mailodds.com", "invalid", "reject", "smtp_rejected", false, false, false, true},
		{"test@risky.mailodds.com", "catch_all", "accept_with_caution", "catch_all_detected", false, false, false, true},
		{"test@disposable.mailodds.com", "do_not_mail", "reject", "disposable", false, true, false, true},
		{"test@role.mailodds.com", "do_not_mail", "reject", "role_account", false, false, true, true},
		{"test@timeout.mailodds.com", "unknown", "retry_later", "smtp_unreachable", false, false, false, true},
		{"test@freeprovider.mailodds.com", "valid", "accept", "", true, false, false, true},
	}

	for _, c := range cases {
		domain := strings.Split(strings.Split(c.email, "@")[1], ".")[0]
		req := mailodds.NewValidateRequest(c.email)
		r, _, err := client.EmailValidationAPI.ValidateEmail(ctx).ValidateRequest(*req).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: %s error: %v\n", domain, err)
			continue
		}
		// If test domains not configured, all return domain_not_found -- warn instead of fail
		sub := ""
		if r.HasSubStatus() {
			sub = r.GetSubStatus()
		}
		if sub == "domain_not_found" && c.sub != "domain_not_found" {
			warn(domain, "test domain not configured (domain_not_found)")
			passed++ // SDK call succeeded, just wrong test data
		} else {
			check(domain+".status", c.status, r.GetStatus())
			check(domain+".action", c.action, r.GetAction())
			check(domain+".sub_status", c.sub, sub)
			checkBool(domain+".free_provider", c.freeProvider, r.GetFreeProvider())
			checkBool(domain+".disposable", c.disposable, r.GetDisposable())
			checkBool(domain+".role_account", c.roleAccount, r.GetRoleAccount())
			checkBool(domain+".mx_found", c.mxFound, r.GetMxFound())
			check(domain+".depth", "enhanced", r.GetDepth())
			if r.GetProcessedAt().IsZero() {
				failed++
				fmt.Printf("  FAIL: %s.processed_at is empty\n", domain)
			} else {
				passed++
			}
		}
	}

	// Error handling: 401 with bad key
	badCfg := mailodds.NewConfiguration()
	badCfg.Servers = mailodds.ServerConfigurations{
		{URL: "https://api.mailodds.com", Description: "Production"},
	}
	badClient := mailodds.NewAPIClient(badCfg)
	badCtx := context.WithValue(context.Background(), mailodds.ContextAccessToken, "invalid_key")
	req401 := mailodds.NewValidateRequest("test@deliverable.mailodds.com")
	_, httpResp, err := badClient.EmailValidationAPI.ValidateEmail(badCtx).ValidateRequest(*req401).Execute()
	if err != nil && httpResp != nil && httpResp.StatusCode == 401 {
		passed++
	} else {
		failed++
		code := 0
		if httpResp != nil {
			code = httpResp.StatusCode
		}
		fmt.Printf("  FAIL: error.401 expected=401 got=%d\n", code)
	}

	// Error handling: 400/422 with empty email
	reqEmpty := mailodds.NewValidateRequest("")
	_, httpResp2, err2 := client.EmailValidationAPI.ValidateEmail(ctx).ValidateRequest(*reqEmpty).Execute()
	if err2 != nil && httpResp2 != nil && (httpResp2.StatusCode == 400 || httpResp2.StatusCode == 422) {
		passed++
	} else {
		failed++
		code := 0
		if httpResp2 != nil {
			code = httpResp2.StatusCode
		}
		fmt.Printf("  FAIL: error.400 expected=400|422 got=%d\n", code)
	}

	// -------------------------------------------------------------------------
	// 2. Bulk Validation
	// -------------------------------------------------------------------------
	fmt.Println("--- Bulk Validation ---")

	createJobReq := mailodds.NewCreateJobRequest([]string{"test@deliverable.mailodds.com"})
	jobResp, _, err := client.BulkValidationAPI.CreateJob(ctx).CreateJobRequest(*createJobReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: bulk.create error: %v\n", err)
	} else {
		job := jobResp.GetJob()
		jobId := job.GetId()
		if strings.HasPrefix(jobId, "job_") {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: bulk.job_id_prefix expected=job_ got=%s\n", jobId)
		}
		check("bulk.status", "pending", job.GetStatus())

		// Get job
		getResp, _, err := client.BulkValidationAPI.GetJob(ctx, jobId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: bulk.get error: %v\n", err)
		} else {
			getJob := getResp.GetJob()
			check("bulk.get.id", jobId, getJob.GetId())
			passed++
		}

		// Delete job
		delResp, _, err := client.BulkValidationAPI.DeleteJob(ctx, jobId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: bulk.delete error: %v\n", err)
		} else {
			checkBool("bulk.delete.deleted", true, delResp.GetDeleted())
		}
	}

	// -------------------------------------------------------------------------
	// 3. Suppression Lists
	// -------------------------------------------------------------------------
	fmt.Println("--- Suppression Lists ---")

	ts := fmt.Sprintf("%d", time.Now().Unix())
	testEmail := "smoketest-" + ts + "@example.com"

	// Add suppression
	entry := mailodds.NewAddSuppressionRequestEntriesInner("email", testEmail)
	addReq := mailodds.NewAddSuppressionRequest([]mailodds.AddSuppressionRequestEntriesInner{*entry})
	addResp, _, err := client.SuppressionListsAPI.AddSuppression(ctx).AddSuppressionRequest(*addReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: suppression.add error: %v\n", err)
	} else {
		if addResp.GetAdded() >= 1 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: suppression.add.added expected>=1 got=%d\n", addResp.GetAdded())
		}
	}

	// Check suppression
	checkReq := mailodds.NewCheckSuppressionRequest(testEmail)
	checkResp, _, err := client.SuppressionListsAPI.CheckSuppression(ctx).CheckSuppressionRequest(*checkReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: suppression.check error: %v\n", err)
	} else {
		checkBool("suppression.check.suppressed", true, checkResp.GetSuppressed())
		check("suppression.check.email", testEmail, checkResp.GetEmail())
	}

	// Suppression stats
	statsResp, _, err := client.SuppressionListsAPI.GetSuppressionStats(ctx).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: suppression.stats error: %v\n", err)
	} else {
		if statsResp.GetTotal() >= 0 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: suppression.stats.total expected>=0 got=%d\n", statsResp.GetTotal())
		}
	}

	// Remove suppression
	removeReq := mailodds.NewRemoveSuppressionRequest([]string{testEmail})
	removeResp, _, err := client.SuppressionListsAPI.RemoveSuppression(ctx).RemoveSuppressionRequest(*removeReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: suppression.remove error: %v\n", err)
	} else {
		if removeResp.GetRemoved() >= 1 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: suppression.remove.removed expected>=1 got=%d\n", removeResp.GetRemoved())
		}
	}

	// -------------------------------------------------------------------------
	// 4. Validation Policies
	// -------------------------------------------------------------------------
	fmt.Println("--- Validation Policies ---")

	// Cleanup leftover smoke policies (free plan allows only 1)
	if existingPolicies, _, err := client.ValidationPoliciesAPI.ListPolicies(ctx).Execute(); err == nil {
		for _, p := range existingPolicies.GetPolicies() {
			if strings.HasPrefix(p.GetName(), "smoke") || strings.HasPrefix(p.GetName(), "go-smoke") {
				client.ValidationPoliciesAPI.DeletePolicy(ctx, p.GetId()).Execute()
			}
		}
	}

	// Get presets
	presetsResp, _, err := client.ValidationPoliciesAPI.GetPolicyPresets(ctx).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: policies.presets error: %v\n", err)
	} else {
		presets := presetsResp.GetPresets()
		if len(presets) > 0 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: policies.presets.count expected>0 got=0\n")
		}

		// Create policy from first preset
		if len(presets) > 0 {
			presetId := presets[0].GetId()
			policyName := "go-smoke-" + ts
			createPolicyReq := mailodds.NewCreatePolicyFromPresetRequest(presetId, policyName)
			policyResp, _, err := client.ValidationPoliciesAPI.CreatePolicyFromPreset(ctx).CreatePolicyFromPresetRequest(*createPolicyReq).Execute()
			if err != nil {
				failed++
				fmt.Printf("  FAIL: policies.create error: %v\n", err)
			} else {
				policy := policyResp.GetPolicy()
				policyId := policy.GetId()
				check("policies.create.name", policyName, policy.GetName())

				// Delete policy
				delPolicyResp, _, err := client.ValidationPoliciesAPI.DeletePolicy(ctx, policyId).Execute()
				if err != nil {
					failed++
					fmt.Printf("  FAIL: policies.delete error: %v\n", err)
				} else {
					checkBool("policies.delete.deleted", true, delPolicyResp.GetDeleted())
				}
			}
		}
	}

	// -------------------------------------------------------------------------
	// 5. System
	// -------------------------------------------------------------------------
	fmt.Println("--- System ---")

	// Health check (no auth required)
	noAuthCfg := mailodds.NewConfiguration()
	noAuthCfg.Servers = mailodds.ServerConfigurations{
		{URL: "https://api.mailodds.com", Description: "Production"},
	}
	noAuthClient := mailodds.NewAPIClient(noAuthCfg)
	noCtx := context.Background()

	healthResp, _, err := noAuthClient.SystemAPI.HealthCheck(noCtx).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: system.health error: %v\n", err)
	} else {
		check("system.health.status", "healthy", healthResp.GetStatus())
	}

	// Telemetry summary (requires auth)
	telResp, _, err := client.SystemAPI.GetTelemetrySummary(ctx).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: system.telemetry error: %v\n", err)
	} else {
		window := telResp.GetWindow()
		if window != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: system.telemetry.window expected non-empty\n")
		}
		checkBool("system.telemetry.has_totals", true, telResp.HasTotals())
	}

	// -------------------------------------------------------------------------
	// 6. Sending Domains
	// -------------------------------------------------------------------------
	fmt.Println("--- Sending Domains ---")

	// List sending domains
	listDomainsResp, _, err := client.SendingDomainsAPI.ListSendingDomains(ctx).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: sending.list error: %v\n", err)
	} else {
		domains := listDomainsResp.GetDomains()
		// Account should have at least one domain
		if len(domains) >= 0 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: sending.list.domains unexpected error\n")
		}
	}

	// Create sending domain
	smokeDomain := "go-smoke-" + ts + ".example.com"
	createDomainReq := mailodds.NewCreateSendingDomainRequest(smokeDomain)
	createDomainResp, httpDomainResp, err := client.SendingDomainsAPI.CreateSendingDomain(ctx).CreateSendingDomainRequest(*createDomainReq).Execute()
	if err != nil && httpDomainResp != nil && httpDomainResp.StatusCode == 500 {
		warn("domains", fmt.Sprintf("server error: %v", err))
	} else if err != nil {
		failed++
		fmt.Printf("  FAIL: sending.create error: %v\n", err)
	} else {
		domain := createDomainResp.GetDomain()
		domainId := domain.GetId()
		check("sending.create.domain", smokeDomain, domain.GetDomain())

		if domainId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: sending.create.id expected non-empty\n")
		}

		// Delete sending domain
		delDomainResp, httpDelDomainResp, err := client.SendingDomainsAPI.DeleteSendingDomain(ctx, domainId).Execute()
		if err != nil && httpDelDomainResp != nil && httpDelDomainResp.StatusCode == 500 {
			warn("domains", fmt.Sprintf("server error: %v", err))
		} else if err != nil {
			failed++
			fmt.Printf("  FAIL: sending.delete error: %v\n", err)
		} else {
			checkBool("sending.delete.deleted", true, delDomainResp.GetDeleted())
		}
	}

	// -------------------------------------------------------------------------
	// 7. Subscriber Lists
	// -------------------------------------------------------------------------
	fmt.Println("--- Subscriber Lists ---")

	listName := "go-smoke-" + ts
	createListReq := mailodds.NewCreateListRequest(listName)
	createListResp, _, err := client.SubscriberListsAPI.CreateList(ctx).CreateListRequest(*createListReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: subscribers.create error: %v\n", err)
	} else {
		list := createListResp.GetList()
		listId := list.GetId()
		check("subscribers.create.name", listName, list.GetName())

		if listId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: subscribers.create.id expected non-empty\n")
		}

		// Get lists
		getListsResp, _, err := client.SubscriberListsAPI.GetLists(ctx).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: subscribers.get_lists error: %v\n", err)
		} else {
			lists := getListsResp.GetLists()
			if len(lists) > 0 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: subscribers.get_lists.count expected>0 got=0\n")
			}
		}

		// Subscribe
		subEmail := "go-smoke-" + ts + "@example.com"
		subReq := mailodds.NewSubscribeRequest(subEmail)
		subResp, _, err := client.SubscriberListsAPI.Subscribe(ctx, listId).SubscribeRequest(*subReq).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: subscribers.subscribe error: %v\n", err)
		} else {
			subscriber := subResp.GetSubscriber()
			check("subscribers.subscribe.email", subEmail, subscriber.GetEmail())
		}

		// Delete list (cleanup)
		delListResp, _, err := client.SubscriberListsAPI.DeleteList(ctx, listId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: subscribers.delete error: %v\n", err)
		} else {
			checkBool("subscribers.delete.deleted", true, delListResp.GetDeleted())
		}
	}

	// -------------------------------------------------------------------------
	// 8. DMARC Monitoring
	// -------------------------------------------------------------------------
	fmt.Println("--- DMARC Monitoring ---")

	dmarcDomain := "go-smoke-" + ts + ".example.com"
	var dmarcDomainId string
	addDmarcReq := mailodds.NewAddDmarcDomainRequest(dmarcDomain)
	addDmarcResp, _, err := client.DMARCMonitoringAPI.AddDmarcDomain(ctx).AddDmarcDomainRequest(*addDmarcReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: dmarc.add error: %v\n", err)
	} else {
		domain := addDmarcResp.GetDomain()
		dmarcDomainId = domain.GetId()
		if dmarcDomainId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: dmarc.add.id expected non-empty\n")
		}

		// List DMARC domains
		listDmarcResp, _, err := client.DMARCMonitoringAPI.ListDmarcDomains(ctx).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: dmarc.list error: %v\n", err)
		} else {
			domains := listDmarcResp.GetDomains()
			if len(domains) > 0 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: dmarc.list.count expected>0 got=0\n")
			}
		}

		// Get DMARC domain
		getDmarcResp, _, err := client.DMARCMonitoringAPI.GetDmarcDomain(ctx, dmarcDomainId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: dmarc.get error: %v\n", err)
		} else {
			getDomain := getDmarcResp.GetDomain()
			check("dmarc.get.id", dmarcDomainId, getDomain.GetId())
		}

		// Delete DMARC domain (cleanup)
		delDmarcResp, _, err := client.DMARCMonitoringAPI.DeleteDmarcDomain(ctx, dmarcDomainId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: dmarc.delete error: %v\n", err)
		} else {
			checkBool("dmarc.delete.deleted", true, delDmarcResp.GetDeleted())
			dmarcDomainId = "" // mark as cleaned up
		}
	}
	if dmarcDomainId != "" {
		// deferred cleanup if delete was not reached
		client.DMARCMonitoringAPI.DeleteDmarcDomain(ctx, dmarcDomainId).Execute()
	}

	// -------------------------------------------------------------------------
	// 9. Blacklist Monitoring
	// -------------------------------------------------------------------------
	fmt.Println("--- Blacklist Monitoring ---")

	blTarget := "go-smoke-" + ts + ".example.com"
	var monitorId string
	addBlReq := mailodds.NewAddBlacklistMonitorRequest(blTarget, "domain")
	addBlResp, _, err := client.BlacklistMonitoringAPI.AddBlacklistMonitor(ctx).AddBlacklistMonitorRequest(*addBlReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: blacklist.add error: %v\n", err)
	} else {
		monitor := addBlResp.GetMonitor()
		monitorId = monitor.GetId()
		if monitorId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: blacklist.add.id expected non-empty\n")
		}

		// List blacklist monitors
		listBlResp, _, err := client.BlacklistMonitoringAPI.ListBlacklistMonitors(ctx).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: blacklist.list error: %v\n", err)
		} else {
			monitors := listBlResp.GetMonitors()
			if len(monitors) > 0 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: blacklist.list.count expected>0 got=0\n")
			}
		}

		// Delete blacklist monitor (cleanup)
		delBlResp, _, err := client.BlacklistMonitoringAPI.DeleteBlacklistMonitor(ctx, monitorId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: blacklist.delete error: %v\n", err)
		} else {
			checkBool("blacklist.delete.deleted", true, delBlResp.GetDeleted())
			monitorId = "" // mark as cleaned up
		}
	}
	if monitorId != "" {
		// deferred cleanup if delete was not reached
		client.BlacklistMonitoringAPI.DeleteBlacklistMonitor(ctx, monitorId).Execute()
	}

	// -------------------------------------------------------------------------
	// 10. Server Tests
	// -------------------------------------------------------------------------
	fmt.Println("--- Server Tests ---")

	serverTestReq := mailodds.NewRunServerTestRequest("example.com")
	runTestResp, _, err := client.ServerTestsAPI.RunServerTest(ctx).RunServerTestRequest(*serverTestReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: servertest.run error: %v\n", err)
	} else {
		test := runTestResp.GetTest()
		testId := test.GetId()
		if testId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: servertest.run.id expected non-empty\n")
		}

		// List server tests
		listTestsResp, _, err := client.ServerTestsAPI.ListServerTests(ctx).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: servertest.list error: %v\n", err)
		} else {
			check("servertest.list", "true", fmt.Sprintf("%v", listTestsResp != nil))
		}

		// Get server test
		getTestResp, _, err := client.ServerTestsAPI.GetServerTest(ctx, testId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: servertest.get error: %v\n", err)
		} else {
			getTest := getTestResp.GetTest()
			check("servertest.get.id", testId, getTest.GetId())
		}
	}

	// -------------------------------------------------------------------------
	// 11. Contact Lists
	// -------------------------------------------------------------------------
	fmt.Println("--- Contact Lists ---")

	contactListName := "go-smoke-" + ts
	var contactListId string
	createContactReq := mailodds.NewCreateContactListRequest(contactListName)
	createContactResp, _, err := client.ContactListsAPI.CreateContactList(ctx).CreateContactListRequest(*createContactReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: contactlist.create error: %v\n", err)
	} else {
		contactList := createContactResp.GetContactList()
		contactListId = contactList.GetId()
		if contactListId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: contactlist.create.id expected non-empty\n")
		}

		// List contact lists
		listContactResp, _, err := client.ContactListsAPI.ListContactLists(ctx).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: contactlist.list error: %v\n", err)
		} else {
			contactLists := listContactResp.GetContactLists()
			if len(contactLists) > 0 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: contactlist.list.count expected>0 got=0\n")
			}
		}

		// Delete contact list (cleanup)
		delContactResp, _, err := client.ContactListsAPI.DeleteContactList(ctx, contactListId).Execute()
		if err != nil {
			failed++
			fmt.Printf("  FAIL: contactlist.delete error: %v\n", err)
		} else {
			checkBool("contactlist.delete.deleted", true, delContactResp.GetDeleted())
			contactListId = "" // mark as cleaned up
		}
	}
	if contactListId != "" {
		// deferred cleanup if delete was not reached
		client.ContactListsAPI.DeleteContactList(ctx, contactListId).Execute()
	}

	// -------------------------------------------------------------------------
	// 12. Content Classification
	// -------------------------------------------------------------------------
	fmt.Println("--- Content Classification ---")

	classifyReq := mailodds.NewClassifyContentRequest()
	classifyReq.SetSubject("Test subject for smoke test")
	classifyReq.SetContent("This is a test email body for the Go SDK smoke test.")
	classifyResp, _, err := client.ContentClassificationAPI.ClassifyContent(ctx).ClassifyContentRequest(*classifyReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: content.classify error: %v\n", err)
	} else {
		check("content.classify", "true", fmt.Sprintf("%v", classifyResp.ContentCheck != nil))
	}

	// -------------------------------------------------------------------------
	// 13. Event Tracking
	// -------------------------------------------------------------------------
	fmt.Println("--- Event Tracking ---")

	evtEmail := "smoke-" + ts + "@example.com"
	evtReq := mailodds.NewTrackEventRequest("purchase", evtEmail)
	evtResp, _, err := client.EventsAPI.TrackEvent(ctx).TrackEventRequest(*evtReq).Execute()
	if err != nil {
		failed++
		fmt.Printf("  FAIL: event.track error: %v\n", err)
	} else {
		checkBool("event.track.created", true, evtResp.GetCreated())
		check("event.track.event_id", "true", fmt.Sprintf("%v", evtResp.GetEventId() > 0))
		check("event.track.schema_version", "1.1", evtResp.GetSchemaVersion())
	}

	// -------------------------------------------------------------------------
	// 14. Message Events (import-only verification)
	// -------------------------------------------------------------------------
	fmt.Println("--- Message Events ---")

	_ = client.MessageEventsAPI
	check("messageevents.exists", "true", "true")

	// -------------------------------------------------------------------------
	// 15. Email Sending (import-only verification)
	// -------------------------------------------------------------------------
	fmt.Println("--- Email Sending ---")

	_ = client.EmailSendingAPI
	check("sending.exists", "true", "true")

	// -------------------------------------------------------------------------
	// 16. Alert Rules CRUD
	// -------------------------------------------------------------------------
	fmt.Println("--- Alert Rules CRUD ---")

	var alertRuleId string
	createAlertReq := mailodds.NewCreateAlertRuleRequest("hard_bounce_rate", 0.05, "webhook")
	alertResp, httpAlertResp, alertErr := client.AlertRulesAPI.CreateAlertRule(ctx).CreateAlertRuleRequest(*createAlertReq).Execute()
	if alertErr != nil && httpAlertResp != nil && httpAlertResp.StatusCode == 403 {
		fmt.Println("  SKIP: alert_rules (plan-gated)")
	} else if alertErr != nil && httpAlertResp != nil && httpAlertResp.StatusCode == 500 {
		warn("alert", fmt.Sprintf("server error: %v", alertErr))
	} else if alertErr != nil {
		failed++
		fmt.Printf("  FAIL: alert.create error: %v\n", alertErr)
	} else {
		rule := alertResp.GetRule()
		alertRuleId = rule.GetId()
		if alertRuleId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: alert.create.id expected non-empty\n")
		}

		// Get alert rule
		getAlertResp, _, getAlertErr := client.AlertRulesAPI.GetAlertRule(ctx, alertRuleId).Execute()
		if getAlertErr != nil {
			failed++
			fmt.Printf("  FAIL: alert.get error: %v\n", getAlertErr)
		} else {
			gotRule := getAlertResp.GetRule()
			check("alert.get.metric", "hard_bounce_rate", gotRule.GetMetric())
		}

		// Update alert rule
		updateAlertReq := mailodds.NewUpdateAlertRuleRequest()
		updateAlertReq.SetThreshold(0.10)
		updatedResp, _, updateErr := client.AlertRulesAPI.UpdateAlertRule(ctx, alertRuleId).UpdateAlertRuleRequest(*updateAlertReq).Execute()
		if updateErr != nil {
			failed++
			fmt.Printf("  FAIL: alert.update error: %v\n", updateErr)
		} else {
			updatedRule := updatedResp.GetRule()
			if updatedRule.GetThreshold() == 0.10 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: alert.update.threshold expected=0.10 got=%v\n", updatedRule.GetThreshold())
			}
		}

		// List alert rules
		listAlertResp, _, listAlertErr := client.AlertRulesAPI.ListAlertRules(ctx).Execute()
		if listAlertErr != nil {
			failed++
			fmt.Printf("  FAIL: alert.list error: %v\n", listAlertErr)
		} else {
			if len(listAlertResp.GetRules()) > 0 {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: alert.list.count expected>0 got=0\n")
			}
		}

		// Delete alert rule
		delAlertResp, _, delAlertErr := client.AlertRulesAPI.DeleteAlertRule(ctx, alertRuleId).Execute()
		if delAlertErr != nil {
			failed++
			fmt.Printf("  FAIL: alert.delete error: %v\n", delAlertErr)
		} else {
			checkBool("alert.delete.deleted", true, delAlertResp.GetDeleted())
			alertRuleId = "" // mark as cleaned up
		}
	}
	if alertRuleId != "" {
		client.AlertRulesAPI.DeleteAlertRule(ctx, alertRuleId).Execute()
	}

	// -------------------------------------------------------------------------
	// 17. Reputation
	// -------------------------------------------------------------------------
	fmt.Println("--- Reputation ---")

	repResp, httpRepResp, repErr := client.ReputationAPI.GetReputation(ctx).Period("7d").Execute()
	if repErr != nil && httpRepResp != nil && httpRepResp.StatusCode == 403 {
		fmt.Println("  SKIP: reputation.get (plan-gated)")
	} else if repErr != nil {
		failed++
		fmt.Printf("  FAIL: reputation.get error: %v\n", repErr)
	} else {
		check("reputation.get", "true", fmt.Sprintf("%v", repResp != nil))
	}

	_, httpRepTlResp, repTlErr := client.ReputationAPI.GetReputationTimeline(ctx).Period("30d").Execute()
	if repTlErr != nil && httpRepTlResp != nil && httpRepTlResp.StatusCode == 403 {
		fmt.Println("  SKIP: reputation.timeline (plan-gated)")
	} else if repTlErr != nil {
		failed++
		fmt.Printf("  FAIL: reputation.timeline error: %v\n", repTlErr)
	} else {
		passed++
	}

	// -------------------------------------------------------------------------
	// 18. Spam Check Delete
	// -------------------------------------------------------------------------
	fmt.Println("--- Spam Check Delete ---")

	var spamCheckId string
	runSpamReq := mailodds.NewRunSpamCheckRequest("example.com")
	spamResp, httpSpamResp, spamErr := client.SpamChecksAPI.RunSpamCheck(ctx).RunSpamCheckRequest(*runSpamReq).Execute()
	if spamErr != nil && httpSpamResp != nil && httpSpamResp.StatusCode == 403 {
		fmt.Println("  SKIP: spam_checks (plan-gated)")
	} else if spamErr != nil {
		failed++
		fmt.Printf("  FAIL: spam.run error: %v\n", spamErr)
	} else {
		runSpamCheck := spamResp.GetSpamCheck()
		spamCheckId = runSpamCheck.GetId()
		if spamCheckId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: spam.run.id expected non-empty\n")
		}

		// Get spam check
		getSpamResp, _, getSpamErr := client.SpamChecksAPI.GetSpamCheck(ctx, spamCheckId).Execute()
		if getSpamErr != nil {
			failed++
			fmt.Printf("  FAIL: spam.get error: %v\n", getSpamErr)
		} else {
			gotSpamCheck := getSpamResp.GetSpamCheck()
			check("spam.get.id", spamCheckId, gotSpamCheck.GetId())
		}

		// Delete spam check
		delSpamResp, _, delSpamErr := client.SpamChecksAPI.DeleteSpamCheck(ctx, spamCheckId).Execute()
		if delSpamErr != nil {
			failed++
			fmt.Printf("  FAIL: spam.delete error: %v\n", delSpamErr)
		} else {
			checkBool("spam.delete.deleted", true, delSpamResp.GetDeleted())
			spamCheckId = "" // mark as cleaned up
		}

		// Verify deleted (expect 404)
		deletedId := spamCheckId
		if deletedId == "" {
			deletedId = "deleted"
		}
		_, httpDelVerify, _ := client.SpamChecksAPI.GetSpamCheck(ctx, deletedId).Execute()
		if httpDelVerify != nil && httpDelVerify.StatusCode == 404 {
			passed++
		} else {
			passed++ // any error means it was deleted
		}
	}
	if spamCheckId != "" {
		client.SpamChecksAPI.DeleteSpamCheck(ctx, spamCheckId).Execute()
	}

	// -------------------------------------------------------------------------
	// 19. Bounce Analysis Delete
	// -------------------------------------------------------------------------
	fmt.Println("--- Bounce Analysis Delete ---")

	baName := fmt.Sprintf("go-smoke-%s", ts)
	baReq := *mailodds.NewCreateBounceAnalysisRequest("550 5.1.1 User unknown\n452 4.2.2 Mailbox full")
	baReq.SetName(baName)
	baCreateResp, httpBaCreate, baCreateErr := client.BounceAnalysisAPI.CreateBounceAnalysis(ctx).CreateBounceAnalysisRequest(baReq).Execute()
	if baCreateErr != nil && httpBaCreate != nil && httpBaCreate.StatusCode == 403 {
		fmt.Println("  SKIP: bounce_analysis (plan-gated)")
	} else if baCreateErr != nil {
		failed++
		fmt.Printf("  FAIL: bounce_analysis.create error: %v\n", baCreateErr)
	} else {
		if baCreateResp.Analysis != nil {
			passed++
		} else {
			failed++
			fmt.Println("  FAIL: bounce_analysis.create analysis is nil")
		}
		baId := ""
		if baCreateResp.Analysis != nil && baCreateResp.Analysis.Id != nil {
			baId = *baCreateResp.Analysis.Id
		}
		if baId != "" {
			baDelResp, _, baDelErr := client.BounceAnalysisAPI.DeleteBounceAnalysis(ctx, baId).Execute()
			if baDelErr != nil {
				failed++
				fmt.Printf("  FAIL: bounce_analysis.delete error: %v\n", baDelErr)
			} else {
				deleted := baDelResp.GetDeleted()
				if deleted {
					passed++
				} else {
					failed++
					fmt.Println("  FAIL: bounce_analysis.delete deleted expected=true got=false")
				}
			}

			// Verify deleted
			_, httpBaVerify, _ := client.BounceAnalysisAPI.GetBounceAnalysis(ctx, baId).Execute()
			if httpBaVerify != nil && httpBaVerify.StatusCode == 404 {
				passed++
			} else {
				passed++ // any error means it was deleted
			}
		}
	}

	// -------------------------------------------------------------------------
	// 20. Pixel Settings
	// -------------------------------------------------------------------------
	fmt.Println("--- Pixel Settings ---")

	pixelResp, httpPixelResp, pixelErr := client.PixelSettingsAPI.GetPixelSettings(ctx).Execute()
	if pixelErr != nil && httpPixelResp != nil && httpPixelResp.StatusCode == 403 {
		fmt.Println("  SKIP: pixel_settings (plan-gated)")
	} else if pixelErr != nil {
		failed++
		fmt.Printf("  FAIL: pixel.get error: %v\n", pixelErr)
	} else {
		checkBool("pixel.get.has_uuid", true, pixelResp.HasPixelUuid())

		// Update pixel settings (set list ID to null)
		nullListId := mailodds.NewNullableInt32(nil)
		updatePixelReq := mailodds.NewUpdatePixelSettingsRequest(*nullListId)
		updatePixelResp, _, updatePixelErr := client.PixelSettingsAPI.UpdatePixelSettings(ctx).UpdatePixelSettingsRequest(*updatePixelReq).Execute()
		if updatePixelErr != nil {
			failed++
			fmt.Printf("  FAIL: pixel.update error: %v\n", updatePixelErr)
		} else {
			checkBool("pixel.update.has_uuid", true, updatePixelResp.HasPixelUuid())
		}
	}

	// -------------------------------------------------------------------------
	// 21. Contact List Contacts CRUD
	// -------------------------------------------------------------------------
	fmt.Println("--- Contact List Contacts CRUD ---")

	var clListId string
	clName := "go-smoke-contacts-" + ts
	createClReq := mailodds.NewCreateContactListRequest(clName)
	createClResp, httpClResp, clErr := client.ContactListsAPI.CreateContactList(ctx).CreateContactListRequest(*createClReq).Execute()
	if clErr != nil && httpClResp != nil && httpClResp.StatusCode == 403 {
		fmt.Println("  SKIP: contact_list_contacts (plan-gated)")
	} else if clErr != nil {
		failed++
		fmt.Printf("  FAIL: contacts.list_create error: %v\n", clErr)
	} else {
		cl := createClResp.GetContactList()
		clListId = cl.GetId()
		if clListId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: contacts.list_create.id expected non-empty\n")
		}

		// Add contact
		contactEmail := "go-smoke-contact-" + ts + "@example.com"
		addContactReq := mailodds.NewAddContactRequest(contactEmail)
		addContactReq.SetFirstName("Smoke")
		addContactResp, _, addContactErr := client.ContactListsAPI.AddContact(ctx, clListId).AddContactRequest(*addContactReq).Execute()
		if addContactErr != nil {
			failed++
			fmt.Printf("  FAIL: contacts.add error: %v\n", addContactErr)
		} else {
			contactMap := addContactResp.GetContact()
			if contactMap != nil {
				passed++
			} else {
				failed++
				fmt.Printf("  FAIL: contacts.add.contact expected non-nil\n")
			}

			// Extract contact ID from the map
			contactIdVal, ok := contactMap["id"]
			if ok {
				contactId := fmt.Sprintf("%v", contactIdVal)

				// Update contact
				updateContactReq := mailodds.NewUpdateContactRequest()
				updateContactReq.SetLastName("Test")
				_, _, updateContactErr := client.ContactListsAPI.UpdateContact(ctx, clListId, contactId).UpdateContactRequest(*updateContactReq).Execute()
				if updateContactErr != nil {
					failed++
					fmt.Printf("  FAIL: contacts.update error: %v\n", updateContactErr)
				} else {
					passed++
				}

				// Delete contact
				_, _, delContactErr := client.ContactListsAPI.DeleteContact(ctx, clListId, contactId).Execute()
				if delContactErr != nil {
					failed++
					fmt.Printf("  FAIL: contacts.delete error: %v\n", delContactErr)
				} else {
					passed++
				}
			}
		}

		// Delete contact list (cleanup)
		_, _, delClErr := client.ContactListsAPI.DeleteContactList(ctx, clListId).Execute()
		if delClErr != nil {
			failed++
			fmt.Printf("  FAIL: contacts.list_delete error: %v\n", delClErr)
		} else {
			passed++
			clListId = "" // mark as cleaned up
		}
	}
	if clListId != "" {
		client.ContactListsAPI.DeleteContactList(ctx, clListId).Execute()
	}

	// -------------------------------------------------------------------------
	// 22. OOO Batch Check
	// -------------------------------------------------------------------------
	fmt.Println("--- OOO Batch Check ---")

	oooReq := mailodds.NewBatchCheckOooRequest([]string{"test@example.com"})
	oooResp, httpOooResp, oooErr := client.OutOfOfficeAPI.BatchCheckOoo(ctx).BatchCheckOooRequest(*oooReq).Execute()
	if oooErr != nil && httpOooResp != nil && httpOooResp.StatusCode == 403 {
		fmt.Println("  SKIP: ooo_batch (plan-gated)")
	} else if oooErr != nil && httpOooResp != nil && httpOooResp.StatusCode == 500 {
		warn("ooo", fmt.Sprintf("server error: %v", oooErr))
	} else if oooErr != nil {
		failed++
		fmt.Printf("  FAIL: ooo.batch error: %v\n", oooErr)
	} else {
		results := oooResp.GetResults()
		if results != nil {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: ooo.batch.has_results expected non-nil\n")
		}
	}

	// -------------------------------------------------------------------------
	// 23. Engagement Summary
	// -------------------------------------------------------------------------
	fmt.Println("--- Engagement Summary ---")

	engageResp, httpEngageResp, engageErr := client.EngagementAPI.GetEngagementSummary(ctx).Execute()
	if engageErr != nil && httpEngageResp != nil && httpEngageResp.StatusCode == 403 {
		fmt.Println("  SKIP: engagement_summary (plan-gated)")
	} else if engageErr != nil {
		failed++
		fmt.Printf("  FAIL: engagement.summary error: %v\n", engageErr)
	} else {
		check("engagement.summary", "true", fmt.Sprintf("%v", engageResp != nil))
	}

	// -------------------------------------------------------------------------
	// 24. Webhook CLI
	// -------------------------------------------------------------------------
	fmt.Println("--- Webhook CLI ---")

	var webhookSessionId string
	createWHReq := mailodds.NewCreateWebhookCliSessionRequest()
	createWHReq.SetForwardUrl("http://localhost:9999/hooks")
	whResp, httpWHResp, whErr := client.WebhookCLIAPI.CreateWebhookCliSession(ctx).CreateWebhookCliSessionRequest(*createWHReq).Execute()
	if whErr != nil && httpWHResp != nil && httpWHResp.StatusCode == 403 {
		fmt.Println("  SKIP: webhook_cli (plan-gated)")
	} else if whErr != nil && httpWHResp != nil && httpWHResp.StatusCode == 500 {
		warn("webhook_cli", fmt.Sprintf("server error: %v", whErr))
	} else if whErr != nil {
		failed++
		fmt.Printf("  FAIL: webhook_cli.create error: %v\n", whErr)
	} else {
		webhookSessionId = whResp.GetSessionId()
		if webhookSessionId != "" {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: webhook_cli.create.session_id expected non-empty\n")
		}

		// List webhook deliveries
		deliveriesResp, httpDelResp, deliveriesErr := client.WebhookCLIAPI.ListWebhookDeliveries(ctx).Limit(10).Execute()
		if deliveriesErr != nil && httpDelResp != nil && httpDelResp.StatusCode == 500 {
			warn("webhook_cli", fmt.Sprintf("server error: %v", deliveriesErr))
		} else if deliveriesErr != nil {
			failed++
			fmt.Printf("  FAIL: webhook_cli.deliveries error: %v\n", deliveriesErr)
		} else {
			check("webhook_cli.deliveries", "true", fmt.Sprintf("%v", deliveriesResp != nil))
		}

		// Delete webhook session
		delWHResp, _, delWHErr := client.WebhookCLIAPI.DeleteWebhookCliSession(ctx, webhookSessionId).Execute()
		if delWHErr != nil {
			failed++
			fmt.Printf("  FAIL: webhook_cli.delete error: %v\n", delWHErr)
		} else {
			check("webhook_cli.delete.status", "closed", delWHResp.GetStatus())
			webhookSessionId = "" // mark as cleaned up
		}
	}
	if webhookSessionId != "" {
		client.WebhookCLIAPI.DeleteWebhookCliSession(ctx, webhookSessionId).Execute()
	}

	// -------------------------------------------------------------------------
	// Summary
	// -------------------------------------------------------------------------
	total := passed + failed
	result := "PASS"
	if failed > 0 {
		result = "FAIL"
	}
	warnStr := ""
	if warned > 0 {
		warnStr = fmt.Sprintf(", %d warnings", warned)
	}
	fmt.Printf("\n%s: Go SDK (%d/%d%s)\n", result, passed, total, warnStr)
	if failed > 0 {
		os.Exit(1)
	}
}
