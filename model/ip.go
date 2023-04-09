package model

import (
	"encoding/json"
	"net/netip"
)

// IPAddr represents a single IP address
//
// Examples:
// IPv4: `10.10.10.1`
// IPv6: `2001:db8::68`
type IPAddr struct {
	addr netip.Addr
}

// ParseIPAddr parses a single IP address from a string
// Supports both IPv4(`10.10.10.1`) and IPv6(`2001:db8::68`)
func ParseIPAddr(s string) (IPAddr, error) {
	addr, err := netip.ParseAddr(s)
	if err != nil {
		return IPAddr{}, err
	}

	return IPAddr{addr}, nil
}

// MustParseIPAddr parses a single IP address from a string
// Supports both IPv4(`10.10.10.1`) and IPv6(`2001:db8::68`)
// It panics on error
func MustParseIPAddr(s string) IPAddr {
	return IPAddr{netip.MustParseAddr(s)}
}

// String returns a single IP address as a string
// `10.10.10.1` for IPv4, `2001:db8::68` for IPv6
func (a IPAddr) String() string {
	return a.addr.String()
}

// MarshalJSON encodes a single IP address as a JSON string
// "10.10.10.1" for IPv4, "2001:db8::68" for IPv6
func (a IPAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.addr.String())
}

// UnmarshalJSON decodes a single IP address from a JSON string
// Supports both IPv4("10.10.10.1") and IPv6("2001:db8::68")
func (a *IPAddr) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*a, err = ParseIPAddr(s)
	return err
}

// IPPrefix represents a network address
//
// Examples:
// IPv4: `10.10.10.0/24`
// IPv6: `2001:db8::/32`
type IPPrefix struct {
	prefix netip.Prefix
}

// ParsePrefix parses a network address from a string
// Supports both IPv4(`10.10.10.0/24`) and IPv6(`2001:db8::/32`)
func ParsePrefix(s string) (IPPrefix, error) {
	p, err := netip.ParsePrefix(s)
	if err != nil {
		return IPPrefix{}, err
	}

	return IPPrefix{p}, nil
}

// MustParsePrefix parses a network address from a string
// Supports both IPv4(`10.10.10.0/24`) and IPv6(`2001:db8::/32`)
// It panics on error
func MustParsePrefix(s string) IPPrefix {
	return IPPrefix{netip.MustParsePrefix(s)}
}

// String returns a network address as a string
// `10.10.10.0/24` for IPv4, `2001:db8::/32` for IPv6
func (p IPPrefix) String() string {
	return p.prefix.String()
}

// MarshalJSON encodes a network address as a JSON string
// "10.10.10.0/24" for IPv4, "2001:db8::/32" for IPv6
func (p IPPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.prefix.String())
}

// UnmarshalJSON decodes a network address from a JSON string
// Supports both IPv4("10.10.10.0/24") and IPv6("2001:db8::/32")
func (p *IPPrefix) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*p, err = ParsePrefix(s)
	return err
}
