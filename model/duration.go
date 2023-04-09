package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Duration represents a time duration
//
// Example: `1w2d3h45m56s67ms`
type Duration struct {
	duration time.Duration
}

var (
	durationRegex    = regexp.MustCompile(`(?:(\d*)w)?(?:(\d*)d)?(?:(\d*)h)?(?:(\d*)m)?(?:(\d*)s)?(?:(\d*)ms)?`)
	durationParts    = [...]time.Duration{time.Hour * 168, time.Hour * 24, time.Hour, time.Minute, time.Second, time.Millisecond}
	durationSuffixes = [...]string{"w", "d", "h", "m", "s", "ms"}
)

// ParseDuration parses a time duration from a string `1w2d3h45m56s67ms`
func ParseDuration(s string) (Duration, error) {
	reMatch := durationRegex.FindAllStringSubmatch(s, -1)

	// should get one and only one match back on the regex
	if len(reMatch) != 1 {
		return Duration{}, errors.New("unexpected regex match")
	}

	var d time.Duration
	for i, match := range reMatch[0] {
		if len(match) == 0 || i == 0 {
			continue
		}

		v, err := strconv.Atoi(match)
		if err != nil {
			return Duration{}, fmt.Errorf("duration parse: %v", err)
		}

		d += time.Duration(v) * durationParts[i-1]
	}

	return Duration{d}, nil
}

// MustParseDuration parses a time duration from a string `1w2d3h45m56s67ms`
// It panics on error
func MustParseDuration(s string) Duration {
	d, err := ParseDuration(s)
	if err != nil {
		panic(err.Error())
	}

	return d
}

// String returns a time duration as a string `1w2d3h45m56s67ms`
func (d Duration) String() string {
	var (
		sb strings.Builder
		n  = d.duration.Nanoseconds()
	)

	for i := range durationParts {
		partNanos := durationParts[i].Nanoseconds()

		if n < partNanos {
			continue
		}

		num := n / partNanos
		n -= num * partNanos
		_, _ = sb.WriteString(fmt.Sprintf("%d%s", n, durationSuffixes[i]))
	}

	return sb.String()
}

// MarshalJSON encodes a time duration as a JSON string "1w2d3h45m56s67ms"
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON decodes a time duration from a JSON string "1w2d3h45m56s67ms"
func (d *Duration) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*d, err = ParseDuration(s)
	return err
}

// DurationWithAuto represents a time duration with an `auto` option
//
// Example: `1w2d3h45m56s67ms`
// Example: `auto`
type DurationWithAuto struct {
	duration Duration
	isAuto   bool
}

// ParseDurationWithAuto parses a time duration with an `auto` option from a string
// Examples of valid strings: `1w2d3h45m56s67ms`, `auto`
func ParseDurationWithAuto(s string) (DurationWithAuto, error) {
	if s == "auto" {
		return DurationWithAuto{isAuto: true}, nil
	}

	d, err := ParseDuration(s)
	if err != nil {
		return DurationWithAuto{}, err
	}

	return DurationWithAuto{duration: d}, nil
}

// MustParseDurationWithAuto parses a time duration with an `auto` option from a string
// Examples of valid strings: `1w2d3h45m56s67ms`, `auto`
// It panics on error
func MustParseDurationWithAuto(s string) DurationWithAuto {
	d, err := ParseDurationWithAuto(s)
	if err != nil {
		panic(err.Error())
	}

	return d
}

// String returns a time duration with an `auto` option as a string
// Examples of valid strings: `1w2d3h45m56s67ms`, `auto`
func (d DurationWithAuto) String() string {
	if d.isAuto {
		return "auto"
	}

	return d.duration.String()
}

// MarshalJSON encodes a time duration with an `auto` option as a JSON string
// Examples of valid JSON strings: "1w2d3h45m56s67ms", "auto"
func (d DurationWithAuto) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON decodes a time duration with an `auto` option from a JSON string
// Examples of valid JSON strings: "1w2d3h45m56s67ms", "auto"
func (d *DurationWithAuto) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*d, err = ParseDurationWithAuto(s)
	return err
}
