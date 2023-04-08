package model

import (
	"encoding/json"
	"net/netip"
)

type IPAddr struct {
	netip.Addr
}

type IPPrefix struct {
	netip.Prefix
}

func ParseIPAddr(s string) (IPAddr, error) {
	addr, err := netip.ParseAddr(s)
	if err != nil {
		return IPAddr{}, err
	}

	return IPAddr{addr}, nil
}

func MustParseIPAddr(s string) IPAddr {
	return IPAddr{netip.MustParseAddr(s)}
}

func (a *IPAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a *IPAddr) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*a, err = ParseIPAddr(s)
	return err
}

func ParsePrefix(s string) (IPPrefix, error) {
	p, err := netip.ParsePrefix(s)
	if err != nil {
		return IPPrefix{}, err
	}

	return IPPrefix{p}, nil
}

func MustParsePrefix(s string) IPPrefix {
	return IPPrefix{netip.MustParsePrefix(s)}
}

func (p *IPPrefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *IPPrefix) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	*p, err = ParsePrefix(s)
	return err
}
