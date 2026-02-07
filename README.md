# MailOdds SDK for Go

Enterprise-ready Go client for the [MailOdds Email Validation API](https://mailodds.com/docs).

## Installation

```bash
go get github.com/mailodds/go-sdk
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"

    mailodds "github.com/mailodds/go-sdk"
)

func main() {
    ctx := context.WithValue(
        context.Background(),
        mailodds.ContextAccessToken,
        "mo_live_your_api_key",
    )

    client := mailodds.NewAPIClient(mailodds.NewConfiguration())
    request := *mailodds.NewValidateRequest("user@example.com")

    result, _, err := client.EmailValidationAPI.ValidateEmail(ctx).
        ValidateRequest(request).Execute()
    if err != nil {
        panic(err)
    }

    // Branch on action for decisioning
    switch result.GetAction() {
    case "accept":
        fmt.Println("Safe to send")
    case "accept_with_caution":
        fmt.Println("Valid but risky -- flag for review")
    case "reject":
        fmt.Println("Do not send")
    case "retry_later":
        fmt.Println("Temporary failure -- retry after backoff")
    }
}
```

## Enterprise Features

This SDK includes enterprise-ready features beyond the generated API client:

### Built-in Retry (429/5xx)

```go
import "github.com/mailodds/go-sdk/enterprise"

cfg := mailodds.NewConfiguration()
cfg.HTTPClient = &http.Client{
    Transport: &enterprise.RetryTransport{
        Base:       http.DefaultTransport,
        MaxRetries: 3,
        BaseDelay:  time.Second,
    },
}

client := mailodds.NewAPIClient(cfg)
```

### Typed Errors

```go
import "github.com/mailodds/go-sdk/enterprise"

result, resp, err := client.EmailValidationAPI.ValidateEmail(ctx).
    ValidateRequest(request).Execute()
if err != nil {
    if resp != nil {
        body, _ := io.ReadAll(resp.Body)
        moErr := enterprise.NewError(resp.StatusCode, body)

        var creditsErr *enterprise.InsufficientCreditsError
        if errors.As(moErr, &creditsErr) {
            fmt.Printf("Need %d credits, have %d\n",
                creditsErr.CreditsNeeded, creditsErr.CreditsAvailable)
            fmt.Printf("Upgrade: %s\n", creditsErr.UpgradeURL)
        }
    }
}
```

### Webhook Signature Verification

```go
import "github.com/mailodds/go-sdk/enterprise"

verifier := enterprise.NewWebhookVerifier("your_webhook_secret")

body, _ := io.ReadAll(r.Body)
signature := r.Header.Get("X-MailOdds-Signature")

event, err := verifier.VerifyAndParse(body, signature)
if err != nil {
    http.Error(w, "Invalid signature", http.StatusUnauthorized)
    return
}
fmt.Printf("Event: %s\n", event["event"])
```

## Response Handling

Branch on the `action` field for decisioning:

| Action | Meaning | Recommended |
|--------|---------|-------------|
| `accept` | Safe to send | Add to mailing list |
| `accept_with_caution` | Valid but risky (catch-all, role account) | Flag for review |
| `reject` | Invalid or disposable | Do not send |
| `retry_later` | Temporary failure | Retry after backoff |

## Test Mode

Use an `mo_test_` prefixed API key with test domains for predictable responses without consuming credits.

## API Reference

- Website: https://mailodds.com
- Full documentation: https://mailodds.com/docs
- OpenAPI spec: https://mailodds.com/openapi.yaml
- All SDKs: https://mailodds.com/sdks

## License

MIT
