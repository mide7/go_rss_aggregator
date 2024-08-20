package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey returns the api key from the headers
// of an http request
// Example:
// Authorization: ApiKey <api_key>
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no api key provided")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("invalid api key")
	}

	return vals[1], nil
}
