package model

import (
	"encoding/json"
	"errors"
)

// ComboMode represent ethernet interface combo mode
// One of: `auto`, `copper`, `sfp`, `sfp-alt`
type ComboMode uint8

const (
	ComboModeAuto ComboMode = iota
	ComboModeCopper
	ComboModeSFP
	ComboModeSFPAlternative
)

const (
	comboModeAutoString           = "auto"
	comboModeCopperString         = "copper"
	comboModeSFPString            = "sfp"
	comboModeSFPAlternativeString = "sfp-alt"
)

// ParseComboMode parses ethernet interface combo mode from a string
// One of: `auto`, `copper`, `sfp`, `sfp-alt`
func ParseComboMode(s string) (ComboMode, error) {
	switch s {
	case comboModeSFPAlternativeString:
		return ComboModeSFPAlternative, nil
	case comboModeSFPString:
		return ComboModeSFP, nil
	case comboModeCopperString:
		return ComboModeCopper, nil
	case comboModeAutoString:
		return ComboModeAuto, nil
	default:
		return 0, errors.New("unknown combo mode")
	}
}

// MustParseComboMode parses ethernet interface combo mode from a string
// One of: `auto`, `copper`, `sfp`, `sfp-alt`
// It panics on error
func MustParseComboMode(s string) ComboMode {
	comboMode, err := ParseComboMode(s)
	if err != nil {
		panic(err.Error())
	}

	return comboMode
}

// String returns ethernet interface combo mode as a string
// One of: `auto`, `copper`, `sfp`, `sfp-alt`
func (cm ComboMode) String() string {
	switch cm {
	case ComboModeSFPAlternative:
		return comboModeSFPAlternativeString
	case ComboModeSFP:
		return comboModeSFPString
	case ComboModeCopper:
		return comboModeCopperString
	default:
		return comboModeAutoString
	}
}

// MarshalJSON encodes ethernet interface combo mode as a JSON string
// One of: "auto", "copper", "sfp", "sfp-alt"
func (cm ComboMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(cm.String())
}

// UnmarshalJSON decodes ethernet interface combo mode from a JSON string
// One of: "auto", "copper", "sfp", "sfp-alt"
func (cm *ComboMode) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*cm, err = ParseComboMode(s)
	return err
}
