package auth

import "testing"

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       map[string][]string
		expectedKey   string
		expectedError error
	}{
		{
			name: "Valid API Key",
			headers: map[string][]string{
				"Authorization": {"ApiKey my-secret-key"},
			},
			expectedKey:   "my-secret-key",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %s, got %s", tt.expectedKey, key)
			}
			if err != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
