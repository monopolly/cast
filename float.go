// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"strconv"
)

func Float(i any) float64 {
	v, _ := ToFloat64E(i)
	return v
}

func Float64(i any) float64 {
	return Float(i)
}

func Float32(i any) float32 {
	v, _ := ToFloat64E(i)
	return float32(v)
}

// ToFloat64E casts an empty interface to a float64.
func ToFloat64E(i any) (float64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case int:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return float64(v), nil
		}
		return 0.0, fmt.Errorf("Unable to Cast %#v to float", i)
	default:
		return 0.0, fmt.Errorf("Unable to Cast %#v to float", i)
	}
}
