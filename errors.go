package go_routeros_client

import (
	"fmt"
	"strings"
)

// Error represents a common RouterOS API error
type Error struct {
	Code    int     `json:"error"`
	Message string  `json:"message"`
	Detail  *string `json:"detail,omitempty"`
}

func (e Error) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("routeros: response code %d, message %q", e.Code, e.Message))

	if e.Detail != nil {
		sb.WriteString(fmt.Sprintf(", detail %q", *e.Detail))
	}

	return sb.String()
}
