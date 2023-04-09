package model

import (
	"encoding/json"
	"errors"
)

// Advertise represent ethernet interface advertise
// One of: `10M-half`, `10M-full`, `100M-half`, `100M-full`, `1000M-half`, `1000M-full`, `2500M-full`, `5000M-full`, `10000M-full`
type Advertise uint8

const (
	AdvertiseNone Advertise = iota
	Advertise10MbpsHalfDuplex
	Advertise10MbpsFullDuplex
	Advertise100MbpsHalfDuplex
	Advertise100MbpsFullDuplex
	Advertise1000MbpsHalfDuplex
	Advertise1000MbpsFullDuplex
	Advertise2500MbpsFullDuplex
	Advertise5000MbpsFullDuplex
	Advertise10000MbpsFullDuplex
)

const (
	advertise10MbpsHalfDuplexString    = "10M-half"
	advertise10MbpsFullDuplexString    = "10M-full"
	advertise100MbpsHalfDuplexString   = "100M-half"
	advertise100MbpsFullDuplexString   = "100M-full"
	advertise1000MbpsHalfDuplexString  = "1000M-half"
	advertise1000MbpsFullDuplexString  = "1000M-full"
	advertise2500MbpsFullDuplexString  = "2500M-full"
	advertise5000MbpsFullDuplexString  = "5000M-full"
	advertise10000MbpsFullDuplexString = "10000M-full"
)

// ParseAdvertise parses ethernet interface advertise from a string
// One of: `10M-half`, `10M-full`, `100M-half`, `100M-full`, `1000M-half`, `1000M-full`, `2500M-full`, `5000M-full`, `10000M-full`
func ParseAdvertise(s string) (Advertise, error) {
	switch s {
	case advertise10000MbpsFullDuplexString:
		return Advertise10000MbpsFullDuplex, nil
	case advertise5000MbpsFullDuplexString:
		return Advertise5000MbpsFullDuplex, nil
	case advertise2500MbpsFullDuplexString:
		return Advertise2500MbpsFullDuplex, nil
	case advertise1000MbpsFullDuplexString:
		return Advertise1000MbpsFullDuplex, nil
	case advertise1000MbpsHalfDuplexString:
		return Advertise1000MbpsHalfDuplex, nil
	case advertise100MbpsFullDuplexString:
		return Advertise100MbpsFullDuplex, nil
	case advertise100MbpsHalfDuplexString:
		return Advertise100MbpsHalfDuplex, nil
	case advertise10MbpsFullDuplexString:
		return Advertise10MbpsFullDuplex, nil
	case advertise10MbpsHalfDuplexString:
		return Advertise10MbpsHalfDuplex, nil
	case "":
		return AdvertiseNone, nil
	default:
		return 0, errors.New("unknown advertise")
	}
}

// MustParseAdvertise parses ethernet interface advertise from a string
// One of: `10M-half`, `10M-full`, `100M-half`, `100M-full`, `1000M-half`, `1000M-full`, `2500M-full`, `5000M-full`, `10000M-full`
// It panics on error
func MustParseAdvertise(s string) Advertise {
	advertise, err := ParseAdvertise(s)
	if err != nil {
		panic(err.Error())
	}

	return advertise
}

// String returns ethernet interface advertise as a string
// One of: `10M-half`, `10M-full`, `100M-half`, `100M-full`, `1000M-half`, `1000M-full`, `2500M-full`, `5000M-full`, `10000M-full`
func (a Advertise) String() string {
	switch a {
	case Advertise10000MbpsFullDuplex:
		return advertise10000MbpsFullDuplexString
	case Advertise5000MbpsFullDuplex:
		return advertise5000MbpsFullDuplexString
	case Advertise2500MbpsFullDuplex:
		return advertise2500MbpsFullDuplexString
	case Advertise1000MbpsFullDuplex:
		return advertise1000MbpsFullDuplexString
	case Advertise1000MbpsHalfDuplex:
		return advertise1000MbpsHalfDuplexString
	case Advertise100MbpsFullDuplex:
		return advertise100MbpsFullDuplexString
	case Advertise100MbpsHalfDuplex:
		return advertise100MbpsHalfDuplexString
	case Advertise10MbpsFullDuplex:
		return advertise10MbpsFullDuplexString
	case Advertise10MbpsHalfDuplex:
		return advertise10MbpsHalfDuplexString
	default:
		return ""
	}
}

// MarshalJSON encodes ethernet interface advertise as a JSON string
// One of: "10M-half", "10M-full", "100M-half", "100M-full", "1000M-half", "1000M-full", "2500M-full", "5000M-full", "10000M-full"
func (a Advertise) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

// UnmarshalJSON decodes ethernet interface advertise from a JSON string
// One of: "10M-half", "10M-full", "100M-half", "100M-full", "1000M-half", "1000M-full", "2500M-full", "5000M-full", "10000M-full"
func (a *Advertise) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*a, err = ParseAdvertise(s)
	return err
}
