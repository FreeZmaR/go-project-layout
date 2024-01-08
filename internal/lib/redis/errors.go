package redis

import (
	"fmt"
	"strings"
)

const (
	jsonTextError   = "json error"
	badRequestError = "bad request"
)

func IsScanJSONError(err error) bool {
	return strings.Contains(err.Error(), jsonTextError)
}

func IsBadRequest(err error) bool {
	return strings.Contains(err.Error(), badRequestError)
}

func newJSONError(err error) error {
	return fmt.Errorf("%s: %w", jsonTextError, err)
}

func newBadRequest(err error) error {
	return fmt.Errorf("%s: %w", badRequestError, err)
}
