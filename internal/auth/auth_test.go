package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  map[string][]string
		expected string
		err      error
	}{
		{
			name: "valid header",
			headers: map[string][]string{
				"Authorization": {"ApiKey valid-api-key"},
			},
			expected: "valid-api-key",
			err:      nil,
		},
		{
			name: "missing Authorization header",
			headers: map[string][]string{
				"Content-Type": {"application/json"},
			},
			expected: "",
			err:      errors.New("no authorization header included"),
		},
		{
			name: "invalid Authorization format",
			headers: map[string][]string{
				"Authorization": {"InvalidFormat"},
			},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
		{
			name: "empty Authorization value",
			headers: map[string][]string{
				"Authorization": {""},
			},
			expected: "",
			err:      errors.New("no authorization header included..."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert map to http.Header
			headers := http.Header(tt.headers)
			apiKey, err := GetAPIKey(headers)
			if apiKey != tt.expected {
				t.Errorf("expected API key %q, got %q", tt.expected, apiKey)
			}
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
