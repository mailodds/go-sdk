// SDK smoke test -- validates build-from-source and API integration.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiURL = "https://api.mailodds.com"

var apiKey string
var passed, failed int

type validateReq struct {
	Email string `json:"email"`
}

type validateResp struct {
	Status   string  `json:"status"`
	Action   string  `json:"action"`
	SubStatus *string `json:"sub_status"`
	TestMode bool    `json:"test_mode"`
}

func check(label, expected, actual string) {
	if expected == actual {
		passed++
	} else {
		failed++
		fmt.Printf("  FAIL: %s expected=%s got=%s\n", label, expected, actual)
	}
}

func apiCall(email, key string) (*validateResp, int, error) {
	body, _ := json.Marshal(validateReq{Email: email})
	req, _ := http.NewRequest("POST", apiURL+"/v1/validate", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, resp.StatusCode, nil
	}

	var result validateResp
	json.Unmarshal(data, &result)
	return &result, resp.StatusCode, nil
}

func main() {
	apiKey = os.Getenv("MAILODDS_TEST_KEY")
	if apiKey == "" {
		fmt.Println("ERROR: MAILODDS_TEST_KEY not set")
		os.Exit(1)
	}

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
		r, _, err := apiCall(c.email, apiKey)
		if err != nil {
			failed++
			fmt.Printf("  FAIL: %s error: %v\n", domain, err)
			continue
		}
		check(domain+".status", c.status, r.Status)
		check(domain+".action", c.action, r.Action)
		sub := ""
		if r.SubStatus != nil {
			sub = *r.SubStatus
		}
		check(domain+".sub_status", c.sub, sub)
		if r.TestMode {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: %s.test_mode expected=true got=false\n", domain)
		}
	}

	// Error handling
	_, code, _ := apiCall("test@deliverable.mailodds.com", "invalid_key")
	check("error.401", "401", fmt.Sprintf("%d", code))

	body, _ := json.Marshal(map[string]string{})
	req, _ := http.NewRequest("POST", apiURL+"/v1/validate", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, _ := client.Do(req)
	if resp != nil {
		resp.Body.Close()
		if resp.StatusCode == 400 || resp.StatusCode == 422 {
			passed++
		} else {
			failed++
			fmt.Printf("  FAIL: error.400 expected=400|422 got=%d\n", resp.StatusCode)
		}
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
