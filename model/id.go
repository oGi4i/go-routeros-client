package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// ID represents the RouterOS identifier
type ID struct {
	value uint64
}

func ParseID(s string) (ID, error) {
	// check if the first char is `*`
	if s[0] != '*' {
		return ID{}, errors.New("id must start with `*`")
	}

	// convert to uint64 from base16
	n, err := strconv.ParseUint(s[1:], 16, 64)
	if err != nil {
		return ID{}, err
	}

	return ID{value: n}, nil
}

func MustParseID(s string) ID {
	id, err := ParseID(s)
	if err != nil {
		panic(err.Error())
	}

	return id
}

func (id *ID) String() string {
	return fmt.Sprintf("*%s", strconv.FormatUint(id.value, 16))
}

func (id *ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *ID) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*id, err = ParseID(s)
	return err
}
