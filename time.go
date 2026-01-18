// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"strings"
	"time"
)

func Time(i any) time.Time {
	v, _ := ToTimeE(i)
	return v
}

func Duration(i any) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

// ToTimeE casts an empty interface to time.Time.
func ToTimeE(i any) (tim time.Time, err error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		d, e := StringToDate(v)
		if e == nil {
			return d, nil
		}
		return time.Time{}, fmt.Errorf("Could not parse Date/Time format: %v\n", e)
	case int:
		return time.Unix(int64(v), 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	default:
		return time.Time{}, fmt.Errorf("Unable to Cast %#v to Time\n", i)
	}
}

// ToDurationE casts an empty interface to time.Duration.
func ToDurationE(i any) (d time.Duration, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int64, int32, int16, int8, int:
		d = time.Duration(Int(s))
		return
	case float32, float64:
		d = time.Duration(Float(s))
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		return
	default:
		err = fmt.Errorf("Unable to Cast %#v to Duration\n", i)
		return
	}
}

// StringToDate casts an empty interface to a time.Time.
func StringToDate(s string) (time.Time, error) {
	return parseDateWith(s, []string{
		time.RFC3339,
		"2006-01-02T15:04:05", // iso8601 without timezone
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC850,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
		"2006-01-02",
		"02 Jan 2006",
		"2006-01-02 15:04:05 -07:00",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05Z07:00", // RFC3339 without T
		"2006-01-02 15:04:05",
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	})
}

func parseDateWith(s string, dates []string) (d time.Time, e error) {
	for _, dateType := range dates {
		if d, e = time.Parse(dateType, s); e == nil {
			return
		}
	}
	return d, fmt.Errorf("Unable to parse date: %s", s)
}
