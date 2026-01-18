// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"strconv"
)

func Int64(i any) int64 {
	v, _ := ToInt64E(i)
	return v
}

func Int(i any) int {
	v, _ := ToIntE(i)
	return v
}

func Int8(i any) int8 {
	v, _ := ToIntE(i)
	return int8(v)
}
func Int16(i any) int16 {
	v, _ := ToIntE(i)
	return int16(v)
}
func Int32(i any) int32 {
	v, _ := ToIntE(i)
	return int32(v)
}

// ToInt64E casts an empty interface to an int64.
func ToInt64E(i any) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return s, nil
	case int:
		return int64(s), nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil

	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int64", i)
	case []byte:
		v, err := strconv.ParseInt(string(s), 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int64", i)
	case float64:
		return int64(s), nil
	case bool:
		if bool(s) {
			return int64(1), nil
		}
		return int64(0), nil
	case nil:
		return int64(0), nil
	default:
		return int64(0), fmt.Errorf("Unable to Cast %#v to int64", i)
	}
}

// ToIntE casts an empty interface to an int.
func ToIntE(i any) (int, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	case []byte:
		v, err := strconv.ParseInt(string(s), 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	case float64:
		return int(s), nil
	case bool:
		if bool(s) {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("Unable to Cast %#v to int", i)
	}
}
