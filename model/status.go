package model

import (
	"encoding/json"
	"errors"
)

// Status represents a boolean status
// One of: `on`, `off`
type Status bool

var (
	StatusOff Status = false
	StatusOn  Status = true
)

const (
	statusOffString = "off"
	statusOnString  = "on"
)

// ParseStatus parses a boolean status from a string
// Valid strings are: `off`, `on`
func ParseStatus(s string) (Status, error) {
	switch s {
	case statusOffString:
		return StatusOff, nil
	case statusOnString:
		return StatusOn, nil
	default:
		return false, errors.New("unknown status")
	}
}

// MustParseStatus parses a boolean status from a string
// Valid strings are: `off`, `on`
// It panics on error
func MustParseStatus(s string) Status {
	t, err := ParseStatus(s)
	if err != nil {
		panic(err.Error())
	}

	return t
}

// String returns a boolean status as a string
// Valid strings are: `off`, `on`
func (s Status) String() string {
	switch s {
	case StatusOn:
		return statusOnString
	default:
		return statusOffString
	}
}

// MarshalJSON encodes a boolean status as a JSON string
// Valid JSON strings are: "off", "on"
func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON decodes a boolean status from a JSON string
// Valid JSON strings are: "off", "on"
func (s *Status) UnmarshalJSON(data []byte) (err error) {
	var str string
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	*s, err = ParseStatus(str)
	return err
}
