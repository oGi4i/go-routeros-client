package model

import (
	"encoding/json"
	"errors"
)

// ARPMode represents an interface ARP mode
// One of: `disabled`, `enabled`, `local-proxy-arp`, `proxy-arp`, `reply-only`
type ARPMode uint8

const (
	ARPModeDisabled ARPMode = iota
	ARPModeEnabled
	ARPModeLocalProxyARP
	ARPModeProxyARP
	ARPModeReplyOnly
)

const (
	arpModeDisabledString      = "disabled"
	arpModeEnabledString       = "enabled"
	arpModeLocalProxyARPString = "local-proxy-arp"
	arpModeProxyARPString      = "proxy-arp"
	arpModeReplyOnlyString     = "reply-only"
)

// ParseARPMode parses ARP mode from a string
// One of: `disabled`, `enabled`, `local-proxy-arp`, `proxy-arp`, `reply-only`
func ParseARPMode(s string) (ARPMode, error) {
	switch s {
	case arpModeReplyOnlyString:
		return ARPModeReplyOnly, nil
	case arpModeProxyARPString:
		return ARPModeProxyARP, nil
	case arpModeLocalProxyARPString:
		return ARPModeLocalProxyARP, nil
	case arpModeEnabledString:
		return ARPModeEnabled, nil
	case arpModeDisabledString:
		return ARPModeDisabled, nil
	default:
		return 0, errors.New("unknown arp mode")
	}
}

// MustParseARPMode parses ARP mode from a string
// One of: `disabled`, `enabled`, `local-proxy-arp`, `proxy-arp`, `reply-only`
// It panics on error
func MustParseARPMode(s string) ARPMode {
	arpMode, err := ParseARPMode(s)
	if err != nil {
		panic(err.Error())
	}

	return arpMode
}

// String returns ARP mode as a string
// One of: `disabled`, `enabled`, `local-proxy-arp`, `proxy-arp`, `reply-only`
func (am ARPMode) String() string {
	switch am {
	case ARPModeReplyOnly:
		return arpModeReplyOnlyString
	case ARPModeProxyARP:
		return arpModeProxyARPString
	case ARPModeLocalProxyARP:
		return arpModeLocalProxyARPString
	case ARPModeEnabled:
		return arpModeEnabledString
	default:
		return arpModeDisabledString
	}
}

// MarshalJSON encodes ARP mode as a JSON string
// One of: "disabled", "enabled", "local-proxy-arp", "proxy-arp", "reply-only"
func (am ARPMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(am.String())
}

// UnmarshalJSON decodes ARP mode from a JSON string
// One of: "disabled", "enabled", "local-proxy-arp", "proxy-arp", "reply-only"
func (am *ARPMode) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*am, err = ParseARPMode(s)
	return err
}
