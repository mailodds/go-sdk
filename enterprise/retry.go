package enterprise

import (
	"bytes"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"time"
)

// RetryTransport wraps an http.RoundTripper with retry logic.
// Retries on 429 (rate limit) and 5xx (server error) with exponential backoff + jitter.
// Does NOT retry on 4xx (except 429).
type RetryTransport struct {
	Base       http.RoundTripper
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
}

// NewRetryTransport creates a RetryTransport with sensible defaults.
func NewRetryTransport(base http.RoundTripper) *RetryTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &RetryTransport{
		Base:       base,
		MaxRetries: 3,
		BaseDelay:  1 * time.Second,
		MaxDelay:   30 * time.Second,
	}
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Buffer the request body so it can be replayed on retries.
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body.Close()
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.MaxRetries; attempt++ {
		// Reset body for each attempt.
		if bodyBytes != nil {
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		} else {
			req.Body = nil
		}

		resp, err = t.Base.RoundTrip(req)
		if err != nil {
			if attempt == t.MaxRetries {
				return nil, err
			}
			t.sleep(attempt)
			continue
		}

		if !shouldRetry(resp.StatusCode) || attempt == t.MaxRetries {
			return resp, nil
		}

		resp.Body.Close()
		t.sleep(attempt)
	}

	return resp, err
}

func shouldRetry(code int) bool {
	return code == 429 || code >= 500
}

func (t *RetryTransport) sleep(attempt int) {
	delay := math.Min(
		float64(t.BaseDelay)*math.Pow(2, float64(attempt)),
		float64(t.MaxDelay),
	)
	jitter := delay * 0.1 * rand.Float64()
	time.Sleep(time.Duration(delay + jitter))
}
