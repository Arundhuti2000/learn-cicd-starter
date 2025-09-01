package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

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

	if err == nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}