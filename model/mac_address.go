package model

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// MACAddress represents a MAC address
//
// Example: 11:22:33:44:55:66
type MACAddress uint64

const (
	macAddressLength = 12

	macAddressSeparator = ":"
	zeroChar            = '0'
)

// ParseMACAddress parses a MAC address from a string `11:22:33:44:55:66`
func ParseMACAddress(s string) (MACAddress, error) {
	sanitized := strings.Replace(s, macAddressSeparator, "", -1)
	if len(sanitized) != macAddressLength {
		return 0, errors.New("invalid mac address")
	}

	// convert to uint64 from base16
	n, err := strconv.ParseUint(sanitized, 16, 64)
	if err != nil {
		return 0, err
	}

	return MACAddress(n), nil
}

// MustParseMACAddress parses a MAC address from a string `11:22:33:44:55:66`
// It panics on error
func MustParseMACAddress(s string) MACAddress {
	macAddress, err := ParseMACAddress(s)
	if err != nil {
		panic(err.Error())
	}

	return macAddress
}

// String returns a MAC address as a string `11:22:33:44:55:66`
func (ma MACAddress) String() string {
	// when used in queries it is important to use uppercase `ABCDEF` instead of `abcdef`
	// otherwise the filter will not match
	asHex := strings.ToUpper(
		strconv.FormatUint(uint64(ma), 16),
	)

	zerosCount := macAddressLength - len(asHex)
	prefix := make([]byte, zerosCount)
	for i := 0; i < zerosCount; i++ {
		prefix[i] = zeroChar
	}

	asHex = string(prefix) + asHex

	return strings.Join(
		[]string{
			asHex[0:2],
			asHex[2:4],
			asHex[4:6],
			asHex[6:8],
			asHex[8:10],
			asHex[10:12],
		},
		macAddressSeparator,
	)
}

// MarshalJSON encodes a MAC address as a JSON string "11:22:33:44:55:66"
func (ma MACAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ma.String())
}

// UnmarshalJSON decodes a MAC address from a JSON string "11:22:33:44:55:66"
func (ma *MACAddress) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*ma, err = ParseMACAddress(s)
	return err
}
