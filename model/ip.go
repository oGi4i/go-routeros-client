package model

import (
	"encoding/json"
	"fmt"
	"net/netip"
	"strings"
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

	if s == "" {
		a = nil
		return nil
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

// Gateway represents a route gateway
//
// Examples:
// `10.10.10.1`
// `10.10.10.1%ether1`
// `ether1`
type Gateway struct {
	addr          *IPAddr
	interfaceName *string
}

// ParseGateway parses a route gateway from a string
// Example valid strings: `10.10.10.1`, `10.10.10.1%ether1`, `ether1`
func ParseGateway(s string) (Gateway, error) {
	// try to parse as IP
	addr, err := ParseIPAddr(s)
	if err != nil {
		separated := strings.Split(s, "%")
		// only interface name case
		if len(separated) < 2 {
			return Gateway{interfaceName: &s}, nil
		}

		// try to parse first part as IP
		addr, err = ParseIPAddr(separated[0])
		if err != nil {
			return Gateway{}, err
		}

		// both IP address and interface name case
		return Gateway{addr: &addr, interfaceName: &separated[1]}, nil
	}

	// only IP address case
	return Gateway{addr: &addr}, nil
}

// MustParseGateway parses a route gateway from a string
// Example valid strings: `10.10.10.1`, `10.10.10.1%ether1`, `ether1`
// It panics on error
func MustParseGateway(s string) Gateway {
	gateway, err := ParseGateway(s)
	if err != nil {
		panic(err.Error())
	}

	return gateway
}

// String returns a route gateway as a string
// Example valid strings: `10.10.10.1`, `10.10.10.1%ether1`, `ether1`
func (g Gateway) String() string {
	switch {
	// IP address
	case g.addr != nil && g.interfaceName == nil:
		return g.addr.String()
	// interface name
	case g.interfaceName != nil && g.addr == nil:
		return *g.interfaceName
	// both IP address and interface name
	default:
		return fmt.Sprintf("%s%%%s", g.addr.String(), *g.interfaceName)
	}
}

// MarshalJSON encodes a route gateway as a JSON string
// Example valid JSON strings: "10.10.10.1", "10.10.10.1%ether1", "ether1"
func (g Gateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// UnmarshalJSON decodes a route gateway from a JSON string
// Example valid JSON strings: "10.10.10.1", "10.10.10.1%ether1", "ether1"
func (g *Gateway) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*g, err = ParseGateway(s)
	return err
}
