package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headerVal string
		wantKey   string
		wantErr   error
	}{
		{
			name:      "Valid ApiKey header",
			headerVal: "ApiKey abc123",
			wantKey:   "abc123",
			wantErr:   nil,
		},
		{
			name:      "Missing Authorization header",
			headerVal: "",
			wantKey:   "",
			wantErr:   ErrNoAuthHeaderIncluded,
		},
		{
			name:      "Malformed header (wrong scheme)",
			headerVal: "Bearer abc123",
			wantKey:   "",
			wantErr:   nil,
		},
		{
			name:      "Malformed header (no token)",
			headerVal: "ApiKey",
			wantKey:   "",
			wantErr:   nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := http.Header{}
			if tc.headerVal != "" {
				h.Set("Authorization", tc.headerVal)
			}
			gotKey, gotErr := GetAPIKey(h)
			if tc.wantErr != nil {
				if gotErr != tc.wantErr {
					t.Fatalf("expected error %v, got %v", tc.wantErr, gotErr)
				}
				return
			}
			if tc.name == "Valid ApiKey header" {
				if gotErr != nil {
					t.Fatalf("expected no error, got %v", gotErr)
				}
				if gotKey != tc.wantKey {
					t.Fatalf("expected key %q, got %q", tc.wantKey, gotKey)
				}
				return
			}

			if gotErr == nil {
				t.Fatalf("expected an error, got nil (key=%q)", gotKey)
			}
		})
	}
}
