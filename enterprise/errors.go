package enterprise

import (
	"encoding/json"
	"fmt"
)

// MailOddsError is the base error type for API responses.
type MailOddsError struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`
	Body       string `json:"-"`
}

func (e *MailOddsError) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("mailodds: %d %s: %s (request_id: %s)", e.StatusCode, e.ErrorCode, e.Message, e.RequestID)
	}
	return fmt.Sprintf("mailodds: %d %s: %s", e.StatusCode, e.ErrorCode, e.Message)
}

// Typed error types for specific HTTP status codes.

type BadRequestError struct{ MailOddsError }
type UnauthorizedError struct{ MailOddsError }
type InsufficientCreditsError struct {
	MailOddsError
	CreditsAvailable int    `json:"credits_available"`
	CreditsNeeded    int    `json:"credits_needed"`
	UpgradeURL       string `json:"upgrade_url"`
}
type ForbiddenError struct{ MailOddsError }
type RateLimitError struct{ MailOddsError }
type ServerError struct{ MailOddsError }

// NewError parses an API error response body and returns a typed error.
func NewError(statusCode int, body []byte) error {
	base := MailOddsError{
		StatusCode: statusCode,
		Body:       string(body),
	}

	var parsed struct {
		Error     string `json:"error"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
		// 402-specific
		CreditsAvailable int    `json:"credits_available"`
		CreditsNeeded    int    `json:"credits_needed"`
		UpgradeURL       string `json:"upgrade_url"`
	}
	if err := json.Unmarshal(body, &parsed); err == nil {
		base.ErrorCode = parsed.Error
		base.Message = parsed.Message
		base.RequestID = parsed.RequestID
	}

	switch statusCode {
	case 400:
		return &BadRequestError{base}
	case 401:
		return &UnauthorizedError{base}
	case 402:
		return &InsufficientCreditsError{
			MailOddsError:    base,
			CreditsAvailable: parsed.CreditsAvailable,
			CreditsNeeded:    parsed.CreditsNeeded,
			UpgradeURL:       parsed.UpgradeURL,
		}
	case 403:
		return &ForbiddenError{base}
	case 429:
		return &RateLimitError{base}
	default:
		if statusCode >= 500 {
			return &ServerError{base}
		}
		return &base
	}
}
