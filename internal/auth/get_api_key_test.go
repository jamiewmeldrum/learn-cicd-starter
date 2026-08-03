package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_Valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if key != "my-secret-key" {
		t.Fatalf("expected %q, got %q", "my-secret-key2", key)
	}
}

func TestGetAPIKey_NoAuthorizationHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}