// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"strconv"
)

func Uint8(i any) uint8 {
	v, _ := ToUint64E(i)
	return uint8(v)
}

func Uint16(i any) uint16 {
	v, _ := ToUint64E(i)
	return uint16(v)
}

func Uint32(i any) uint32 {
	v, _ := ToUint64E(i)
	return uint32(v)
}

func Uint64(i any) uint64 {
	v, _ := ToUint64E(i)
	return uint64(v)
}

func Uint(i any) uint {
	return uint(Uint64(i))
}

// ToInt64E casts an empty interface to an int64.
func ToUint64E(i any) (uint64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int64:
		return uint64(s), nil
	case int:
		return uint64(s), nil
	case int32:
		return uint64(s), nil
	case int16:
		return uint64(s), nil
	case int8:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint:
		return uint64(s), nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil

	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int64", i)
	case []byte:
		v, err := strconv.ParseUint(string(s), 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("Unable to Cast %#v to int64", i)
	case float64:
		return uint64(s), nil
	case bool:
		if bool(s) {
			return uint64(1), nil
		}
		return uint64(0), nil
	case nil:
		return uint64(0), nil
	default:
		return uint64(0), fmt.Errorf("Unable to Cast %#v to int64", i)
	}
}
