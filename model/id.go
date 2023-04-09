package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// ID represents a RouterOS identifier
//
// Example: `*ffff1234`
type ID uint64

// ParseID parses a RouterOS identifier from a hex string `*ffffffff`
func ParseID(s string) (ID, error) {
	// check if the first char is `*`
	if s[0] != '*' {
		return 0, errors.New("id must start with `*`")
	}

	// convert to uint64 from base16
	n, err := strconv.ParseUint(s[1:], 16, 64)
	if err != nil {
		return 0, err
	}

	return ID(n), nil
}

// MustParseID parses a RouterOS identifier from a hex string `*ffffffff`
// It panics on error
func MustParseID(s string) ID {
	id, err := ParseID(s)
	if err != nil {
		panic(err.Error())
	}

	return id
}

// String returns a RouterOS identifier as a hex string `*ffffffff`
func (id ID) String() string {
	return fmt.Sprintf("*%s", strconv.FormatUint(uint64(id), 16))
}

// MarshalJSON encodes a RouterOS identifier as a JSON string "*ffffffff"
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON decodes a RouterOS identifier from a JSON string "*ffffffff"
func (id *ID) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*id, err = ParseID(s)
	return err
}
