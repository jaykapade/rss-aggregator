package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts the API key from the Authorization header
//
// The header format should be: "Authorization: ApiKey {api_key}"
//
// Example:
//
//	Authorization: ApiKey 1234567890
//
// Returns:
//
//	string: The extracted API key if successful
//	error: An error if the header is missing or malformed
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return vals[1], nil
}
