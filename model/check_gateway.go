package model

import (
	"encoding/json"
	"errors"
)

// CheckGatewayMode represents an IP route check gateway mode
// One of: `none`, `arp`, `bfd`, `bfd-multihop`, `ping`
type CheckGatewayMode uint8

const (
	CheckGatewayModeNone CheckGatewayMode = iota
	CheckGatewayModeARP
	CheckGatewayModeBFD
	CheckGatewayModeBFDMultiHop
	CheckGatewayModePing
)

const (
	checkGatewayModeNoneString        = "none"
	checkGatewayModeARPString         = "arp"
	checkGatewayModeBFDString         = "bfd"
	checkGatewayModeBFDMultiHopString = "bfd-multihop"
	checkGatewayModePingString        = "ping"
)

// ParseCheckGatewayMode parses IP route check gateway mode from a string
// One of: `none`, `arp`, `bfd`, `bfd-multihop`, `ping`
func ParseCheckGatewayMode(s string) (CheckGatewayMode, error) {
	switch s {
	case checkGatewayModePingString:
		return CheckGatewayModePing, nil
	case checkGatewayModeBFDMultiHopString:
		return CheckGatewayModeBFDMultiHop, nil
	case checkGatewayModeBFDString:
		return CheckGatewayModeBFD, nil
	case checkGatewayModeARPString:
		return CheckGatewayModeARP, nil
	case checkGatewayModeNoneString:
		return CheckGatewayModeNone, nil
	default:
		return 0, errors.New("unknown check gateway mode")
	}
}

// MustParseCheckGatewayMode parses IP route check gateway mode from a string
// One of: `none`, `arp`, `bfd`, `bfd-multihop`, `ping`
// It panics on error
func MustParseCheckGatewayMode(s string) CheckGatewayMode {
	mode, err := ParseCheckGatewayMode(s)
	if err != nil {
		panic(err.Error())
	}

	return mode
}

// String returns IP route check gateway mode as a string
// One of: `none`, `arp`, `bfd`, `bfd-multihop`, `ping`
func (cgm CheckGatewayMode) String() string {
	switch cgm {
	case CheckGatewayModePing:
		return checkGatewayModePingString
	case CheckGatewayModeBFDMultiHop:
		return checkGatewayModeBFDMultiHopString
	case CheckGatewayModeBFD:
		return checkGatewayModeBFDString
	case CheckGatewayModeARP:
		return checkGatewayModeARPString
	default:
		return checkGatewayModeNoneString
	}
}

// MarshalJSON encodes ARP mode as a JSON string
// One of: "none", "arp", "bfd", "bfd-multihop", "ping"
func (cgm CheckGatewayMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(cgm.String())
}

// UnmarshalJSON decodes ARP mode from a JSON string
// One of: "none", "arp", "bfd", "bfd-multihop", "ping"
func (cgm *CheckGatewayMode) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*cgm, err = ParseCheckGatewayMode(s)
	return err
}
