package enterprise

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// WebhookVerifier verifies MailOdds webhook signatures using HMAC-SHA256.
type WebhookVerifier struct {
	secret []byte
}

// NewWebhookVerifier creates a verifier with the given signing secret.
func NewWebhookVerifier(secret string) *WebhookVerifier {
	return &WebhookVerifier{secret: []byte(secret)}
}

// Verify checks if the signature matches the payload.
func (v *WebhookVerifier) Verify(payload []byte, signature string) bool {
	mac := hmac.New(sha256.New, v.secret)
	mac.Write(payload)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(signature))
}

// ErrInvalidSignature is returned when webhook signature verification fails.
var ErrInvalidSignature = errors.New("mailodds: invalid webhook signature")

// VerifyOrError checks the signature and returns an error if invalid.
func (v *WebhookVerifier) VerifyOrError(payload []byte, signature string) error {
	if !v.Verify(payload, signature) {
		return ErrInvalidSignature
	}
	return nil
}

// VerifyAndParse checks the signature and parses the JSON payload into a map.
// Returns ErrInvalidSignature if the signature is invalid.
func (v *WebhookVerifier) VerifyAndParse(payload []byte, signature string) (map[string]interface{}, error) {
	if !v.Verify(payload, signature) {
		return nil, ErrInvalidSignature
	}
	var result map[string]interface{}
	if err := json.Unmarshal(payload, &result); err != nil {
		return nil, err
	}
	return result, nil
}
