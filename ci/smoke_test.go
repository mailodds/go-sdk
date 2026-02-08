// SDK smoke test -- validates build-from-source and API integration using the SDK client.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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

	type tc struct {
		email, status, action, sub string
	}
	cases := []tc{
		{"test@deliverable.mailodds.com", "valid", "accept", ""},
		{"test@invalid.mailodds.com", "invalid", "reject", "smtp_rejected"},
		{"test@risky.mailodds.com", "catch_all", "accept_with_caution", "catch_all_detected"},
		{"test@disposable.mailodds.com", "do_not_mail", "reject", "disposable"},
		{"test@role.mailodds.com", "do_not_mail", "reject", "role_account"},
		{"test@timeout.mailodds.com", "unknown", "retry_later", "smtp_unreachable"},
		{"test@freeprovider.mailodds.com", "valid", "accept", ""},
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
