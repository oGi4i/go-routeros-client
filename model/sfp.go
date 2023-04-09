package model

import (
	"encoding/json"
	"errors"
)

// SFPRateSelectMode represents an interface SFP rate select mode
// One of: `low`, `high`
type SFPRateSelectMode uint8

const (
	SFPRateSelectModeLow SFPRateSelectMode = iota
	SFPRateSelectModeHigh
)

const (
	sfpRateSelectModeLowString  = "low"
	sfpRateSelectModeHighString = "high"
)

// ParseSFPRateSelectMode parses SFP rate select mode from a string
// One of: `low`, `high`
func ParseSFPRateSelectMode(s string) (SFPRateSelectMode, error) {
	switch s {
	case sfpRateSelectModeHighString:
		return SFPRateSelectModeHigh, nil
	case sfpRateSelectModeLowString:
		return SFPRateSelectModeLow, nil
	default:
		return 0, errors.New("unknown sfp rate select mode")
	}
}

// MustParseSFPRateSelectMode parses SFP rate select mode from a string
// One of: `low`, `high`
// It panics on error
func MustParseSFPRateSelectMode(s string) SFPRateSelectMode {
	mode, err := ParseSFPRateSelectMode(s)
	if err != nil {
		panic(err.Error())
	}

	return mode
}

// String returns SFP rate select mode as a string
// One of: `low`, `high`
func (srsm SFPRateSelectMode) String() string {
	switch srsm {
	case SFPRateSelectModeHigh:
		return sfpRateSelectModeHighString
	default:
		return sfpRateSelectModeLowString
	}
}

// MarshalJSON encodes SFP rate select mode as a JSON string
// One of: "low", "high"
func (srsm SFPRateSelectMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(srsm.String())
}

// UnmarshalJSON decodes SFP rate select mode from a JSON string
// One of: "low", "high"
func (srsm *SFPRateSelectMode) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*srsm, err = ParseSFPRateSelectMode(s)
	return err
}
