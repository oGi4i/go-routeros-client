package model

import (
	"encoding/json"
	"errors"
)

// CableSettings represent ethernet interface cable settings
// One of: `default`, `short`, `standard`
type CableSettings uint8

const (
	CableSettingsDefault CableSettings = iota
	CableSettingsShort
	CableSettingsStandard
)

const (
	cableSettingsDefaultString  = "default"
	cableSettingsShortString    = "short"
	cableSettingsStandardString = "standard"
)

// ParseCableSettings parses ethernet interface cable settings from a string
// One of: `default`, `short`, `standard`
func ParseCableSettings(s string) (CableSettings, error) {
	switch s {
	case cableSettingsStandardString:
		return CableSettingsStandard, nil
	case cableSettingsShortString:
		return CableSettingsShort, nil
	case cableSettingsDefaultString:
		return CableSettingsDefault, nil
	default:
		return 0, errors.New("unknown cable settings")
	}
}

// MustParseCableSettings parses ethernet interface cable settings from a string
// One of: `default`, `short`, `standard`
// It panics on error
func MustParseCableSettings(s string) CableSettings {
	cableSettings, err := ParseCableSettings(s)
	if err != nil {
		panic(err.Error())
	}

	return cableSettings
}

// String returns ethernet interface cable settings as a string
// One of: `default`, `short`, `standard`
func (cs CableSettings) String() string {
	switch cs {
	case CableSettingsStandard:
		return cableSettingsStandardString
	case CableSettingsShort:
		return cableSettingsShortString
	default:
		return cableSettingsDefaultString
	}
}

// MarshalJSON encodes ethernet interface cable settings as a JSON string
// One of: "default", "short", "standard"
func (cs CableSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(cs.String())
}

// UnmarshalJSON decodes ethernet interface cable settings from a JSON string
// One of: "default", "short", "standard"
func (cs *CableSettings) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*cs, err = ParseCableSettings(s)
	return err
}
