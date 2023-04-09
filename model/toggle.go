package model

import (
	"encoding/json"
	"errors"
)

// ToggleWithDefault represents a boolean toggle with a `default` option
// One of: `default`, `off`, `on`
type ToggleWithDefault struct {
	value     bool
	isDefault bool
}

var (
	ToggleWithDefaultDefault = ToggleWithDefault{isDefault: true}
	ToggleWithDefaultOff     = ToggleWithDefault{}
	ToggleWithDefaultOn      = ToggleWithDefault{value: true}
)

const (
	toggleWithDefaultDefaultString = "default"
	toggleWithDefaultOffString     = "off"
	toggleWithDefaultOnString      = "on"
)

// ParseToggleWithDefault parses a boolean toggle with a `default` option from a string
// Valid strings are: `default`, `off`, `on`
func ParseToggleWithDefault(s string) (ToggleWithDefault, error) {
	switch s {
	case toggleWithDefaultDefaultString:
		return ToggleWithDefaultDefault, nil
	case toggleWithDefaultOffString:
		return ToggleWithDefaultOff, nil
	case toggleWithDefaultOnString:
		return ToggleWithDefaultOn, nil
	default:
		return ToggleWithDefault{}, errors.New("unknown toggle with default")
	}
}

// MustParseToggleWithDefault parses a boolean toggle with a `default` option from a string
// Valid strings are: `default`, `off`, `on`
// It panics on error
func MustParseToggleWithDefault(s string) ToggleWithDefault {
	t, err := ParseToggleWithDefault(s)
	if err != nil {
		panic(err.Error())
	}

	return t
}

// String returns a boolean toggle with `default` option as a string
// Valid strings are: `default`, `off`, `on`
func (t ToggleWithDefault) String() string {
	switch t {
	case ToggleWithDefaultOn:
		return toggleWithDefaultOnString
	case ToggleWithDefaultOff:
		return toggleWithDefaultOffString
	default:
		return toggleWithDefaultDefaultString
	}
}

// MarshalJSON encodes a boolean toggle with a `default` option as a JSON string
// Valid JSON strings are: "default", "off", "on"
func (t ToggleWithDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON decodes a boolean toggle with a `default` option from a JSON string
// Valid JSON strings are: "default", "off", "on"
func (t *ToggleWithDefault) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t, err = ParseToggleWithDefault(s)
	return err
}

// ToggleWithAuto represents a boolean toggle with a `auto` option
// One of: `auto`, `off`, `on`
type ToggleWithAuto struct {
	value  bool
	isAuto bool
}

var (
	ToggleWithAutoAuto = ToggleWithAuto{isAuto: true}
	ToggleWithAutoOff  = ToggleWithAuto{}
	ToggleWithAutoOn   = ToggleWithAuto{value: true}
)

const (
	toggleWithAutoAutoString = "auto"
	toggleWithAutoOffString  = "off"
	toggleWithAutoOnString   = "on"
)

// ParseToggleWithAuto parses a boolean toggle with a `auto` option from a string
// Valid strings are: `auto`, `off`, `on`
func ParseToggleWithAuto(s string) (ToggleWithAuto, error) {
	switch s {
	case toggleWithAutoAutoString:
		return ToggleWithAutoAuto, nil
	case toggleWithAutoOffString:
		return ToggleWithAutoOff, nil
	case toggleWithAutoOnString:
		return ToggleWithAutoOn, nil
	default:
		return ToggleWithAuto{}, errors.New("unknown toggle with auto")
	}
}

// MustParseToggle parses a boolean toggle with a `auto` option from a string
// Valid strings are: `auto`, `off`, `on`
// It panics on error
func MustParseToggle(s string) ToggleWithAuto {
	t, err := ParseToggleWithAuto(s)
	if err != nil {
		panic(err.Error())
	}

	return t
}

// String returns a boolean toggle with `auto` option as a string
// Valid strings are: `auto`, `off`, `on`
func (t ToggleWithAuto) String() string {
	switch t {
	case ToggleWithAutoOn:
		return toggleWithAutoOnString
	case ToggleWithAutoOff:
		return toggleWithAutoOffString
	default:
		return toggleWithAutoAutoString
	}
}

// MarshalJSON encodes a boolean toggle with a `auto` option as a JSON string
// Valid JSON strings are: "auto", "off", "on"
func (t ToggleWithAuto) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON decodes a boolean toggle with a `auto` option from a JSON string
// Valid JSON strings are: "auto", "off", "on"
func (t *ToggleWithAuto) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*t, err = ParseToggleWithAuto(s)
	return err
}
