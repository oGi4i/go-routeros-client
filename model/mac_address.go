package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type MACAddress struct {
	value uint64
}

func ParseMACAddress(s string) (MACAddress, error) {
	sanitized := strings.Replace(s, ":", "", -1)

	// convert to uint64 from base16
	n, err := strconv.ParseUint(sanitized, 16, 64)
	if err != nil {
		return MACAddress{}, err
	}

	return MACAddress{value: n}, nil
}

func MustParseMACAddress(s string) MACAddress {
	macAddress, err := ParseMACAddress(s)
	if err != nil {
		panic(err.Error())
	}

	return macAddress
}

func (ma *MACAddress) String() string {
	asHex := strconv.FormatUint(ma.value, 16)
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		asHex[0:2],
		asHex[2:4],
		asHex[4:6],
		asHex[6:8],
		asHex[8:10],
		asHex[10:12],
	)
}

func (ma *MACAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ma.value)
}

func (ma *MACAddress) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*ma, err = ParseMACAddress(s)
	return err
}
