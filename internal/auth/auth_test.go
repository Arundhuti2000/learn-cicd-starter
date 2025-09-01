package auth

import (
	"net/http"
	"testing"
)
func TestGetAPIKey(h *testing.T){
	headers := http.Header{}
	_,err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		h.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
	if err == nil || err.Error() != "malformed authorization header" {
		h.Fatalf("expected: malformed authorization header, got: %v", err)
	}
}
func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	got, err := GetAPIKey(headers)
	want := "my-secret-key"

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}