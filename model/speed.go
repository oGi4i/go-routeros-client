package model

import (
	"encoding/json"
	"errors"
)

// Speed represent ethernet interface speed
// One of: `10Mbps`, `100Mbps`, `1Gbps`, `2.5Gbps`, `5Gbps`, `10Gbps`, `25Gbps`, `40Gbps`, `50Gbps`, `100Gbps`, `200Gbps`, `400Gbps`
type Speed uint8

const (
	Speed10Mbps Speed = iota
	Speed100Mbps
	Speed1Gbps
	Speed2AndHalfGbps
	Speed5Gbps
	Speed10Gbps
	Speed25Gbps
	Speed40Gbps
	Speed50Gbps
	Speed100Gbps
	Speed200Gbps
	Speed400Gbps
)

const (
	speed10MbpsString       = "10Mbps"
	speed100MbpsString      = "100Mbps"
	speed1GbpsString        = "1Gbps"
	speed2AndHalfGbpsString = "2.5Gbps"
	speed5GbpsString        = "5Gbps"
	speed10GbpsString       = "10Gbps"
	speed25GbpsString       = "25Gbps"
	speed40GbpsString       = "40Gbps"
	speed50GbpsString       = "50Gbps"
	speed100GbpsString      = "100Gbps"
	speed200GbpsString      = "200Gbps"
	speed400GbpsString      = "400Gbps"
)

// ParseSpeed parses ethernet interface speed from a string
// One of: `10Mbps`, `100Mbps`, `1Gbps`, `2.5Gbps`, `5Gbps`, `10Gbps`, `25Gbps`, `40Gbps`, `50Gbps`, `100Gbps`, `200Gbps`, `400Gbps`
func ParseSpeed(s string) (Speed, error) {
	switch s {
	case speed400GbpsString:
		return Speed400Gbps, nil
	case speed200GbpsString:
		return Speed200Gbps, nil
	case speed100GbpsString:
		return Speed100Gbps, nil
	case speed50GbpsString:
		return Speed50Gbps, nil
	case speed40GbpsString:
		return Speed40Gbps, nil
	case speed25GbpsString:
		return Speed25Gbps, nil
	case speed10GbpsString:
		return Speed10Gbps, nil
	case speed5GbpsString:
		return Speed5Gbps, nil
	case speed2AndHalfGbpsString:
		return Speed2AndHalfGbps, nil
	case speed1GbpsString:
		return Speed1Gbps, nil
	case speed100MbpsString:
		return Speed100Mbps, nil
	case speed10MbpsString:
		return Speed10Mbps, nil
	default:
		return 0, errors.New("unknown speed")
	}
}

// MustParseSpeed parses ethernet interface cable settings from a string
// One of: `10Mbps`, `100Mbps`, `1Gbps`, `2.5Gbps`, `5Gbps`, `10Gbps`, `25Gbps`, `40Gbps`, `50Gbps`, `100Gbps`, `200Gbps`, `400Gbps`
// It panics on error
func MustParseSpeed(s string) Speed {
	speed, err := ParseSpeed(s)
	if err != nil {
		panic(err.Error())
	}

	return speed
}

// String returns ethernet interface speed as a string
// One of: `10Mbps`, `100Mbps`, `1Gbps`, `2.5Gbps`, `5Gbps`, `10Gbps`, `25Gbps`, `40Gbps`, `50Gbps`, `100Gbps`, `200Gbps`, `400Gbps`
func (s Speed) String() string {
	switch s {
	case Speed400Gbps:
		return speed400GbpsString
	case Speed200Gbps:
		return speed200GbpsString
	case Speed100Gbps:
		return speed100GbpsString
	case Speed50Gbps:
		return speed50GbpsString
	case Speed40Gbps:
		return speed40GbpsString
	case Speed25Gbps:
		return speed25GbpsString
	case Speed10Gbps:
		return speed10GbpsString
	case Speed5Gbps:
		return speed5GbpsString
	case Speed2AndHalfGbps:
		return speed2AndHalfGbpsString
	case Speed1Gbps:
		return speed1GbpsString
	case Speed100Mbps:
		return speed100MbpsString
	default:
		return speed10MbpsString
	}
}

// MarshalJSON encodes ethernet interface speed as a JSON string
// One of: "10Mbps", "100Mbps", "1Gbps", "2.5Gbps", "5Gbps", "10Gbps", "25Gbps", "40Gbps", "50Gbps", "100Gbps", "200Gbps", "400Gbps"
func (s Speed) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON decodes ethernet interface speed from a JSON string
// One of: "10Mbps", "100Mbps", "1Gbps", "2.5Gbps", "5Gbps", "10Gbps", "25Gbps", "40Gbps", "50Gbps", "100Gbps", "200Gbps", "400Gbps"
func (s *Speed) UnmarshalJSON(data []byte) (err error) {
	var str string
	if err = json.Unmarshal(data, &str); err != nil {
		return err
	}

	*s, err = ParseSpeed(str)
	return err
}
