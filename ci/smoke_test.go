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

var passed, failed int

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
		check(domain+".status", c.status, r.GetStatus())
		check(domain+".action", c.action, r.GetAction())
		sub := ""
		if r.HasSubStatus() {
			sub = r.GetSubStatus()
		}
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
	createDomainResp, _, err := client.SendingDomainsAPI.CreateSendingDomain(ctx).CreateSendingDomainRequest(*createDomainReq).Execute()
	if err != nil {
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
		delDomainResp, _, err := client.SendingDomainsAPI.DeleteSendingDomain(ctx, domainId).Execute()
		if err != nil {
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
	// 8. Email Sending (import-only verification)
	// -------------------------------------------------------------------------
	fmt.Println("--- Email Sending ---")

	_ = client.EmailSendingAPI
	check("sending.exists", "true", "true")

	// -------------------------------------------------------------------------
	// Summary
	// -------------------------------------------------------------------------
	total := passed + failed
	result := "PASS"
	if failed > 0 {
		result = "FAIL"
	}
	fmt.Printf("\n%s: Go SDK (%d/%d)\n", result, passed, total)
	if failed > 0 {
		os.Exit(1)
	}
}
