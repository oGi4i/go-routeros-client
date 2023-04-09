package model

import (
	"encoding/json"
	"errors"
)

// FECMode represent ethernet interface FEC mode
// One of: `auto`, `fec74`, `fec91`, `off`
type FECMode uint8

const (
	FECModeAuto FECMode = iota
	FECMode74
	FECMode91
	FECModeOff
)

const (
	fecModeAutoString = "auto"
	fecMode74String   = "fec74"
	fecMode91String   = "fec91"
	fecModeOffString  = "off"
)

// ParseFECMode parses ethernet interface FEC mode from a string
// One of: `auto`, `fec74`, `fec91`, `off`
func ParseFECMode(s string) (FECMode, error) {
	switch s {
	case fecModeOffString:
		return FECModeOff, nil
	case fecMode91String:
		return FECMode91, nil
	case fecMode74String:
		return FECMode74, nil
	case fecModeAutoString:
		return FECModeAuto, nil
	default:
		return 0, errors.New("unknown fec mode")
	}
}

// MustParseFECMode parses ethernet interface FEC mode from a string
// One of: `auto`, `fec74`, `fec91`, `off`
// It panics on error
func MustParseFECMode(s string) ComboMode {
	comboMode, err := ParseComboMode(s)
	if err != nil {
		panic(err.Error())
	}

	return comboMode
}

// String returns ethernet interface FEC mode as a string
// One of: `auto`, `fec74`, `fec91`, `off`
func (fm FECMode) String() string {
	switch fm {
	case FECModeOff:
		return fecModeOffString
	case FECMode91:
		return fecMode91String
	case FECMode74:
		return fecMode74String
	default:
		return fecModeAutoString
	}
}

// MarshalJSON encodes ethernet interface FEC mode as a JSON string
// One of: "auto", "fec74", "fec91", "off"
func (fm FECMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(fm.String())
}

// UnmarshalJSON decodes ethernet interface FEC mode from a JSON string
// One of: "auto", "fec74", "fec91", "off"
func (fm *FECMode) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*fm, err = ParseFECMode(s)
	return err
}
