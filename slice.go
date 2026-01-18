// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"reflect"
	"strings"
)

func SliceInterface(i any) []any {
	v, _ := ToSliceE(i)
	return v
}

func SliceBool(i any) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

func SliceString(i any) []string {
	v, _ := ToStringSliceE(i)
	return v
}

func SliceInt(i any) []int {
	v, _ := ToIntSliceE(i)
	return v
}

func SliceInt8(i any) (res []int8) {
	v, _ := ToIntSliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]int8, len(v))
	for i, x := range v {
		res[i] = int8(x)
	}
	return
}

func SliceInt16(i any) (res []int16) {
	v, _ := ToIntSliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]int16, len(v))
	for i, x := range v {
		res[i] = int16(x)
	}
	return
}
func SliceInt32(i any) (res []int32) {
	v, _ := ToIntSliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]int32, len(v))
	for i, x := range v {
		res[i] = int32(x)
	}
	return
}

func SliceInt64(i any) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}

func SliceUint8(i any) (res []uint8) {
	v, _ := ToInt64SliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]uint8, len(v))
	for i, x := range v {
		res[i] = uint8(x)
	}
	return
}
func SliceUint16(i any) (res []uint16) {
	v, _ := ToInt64SliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]uint16, len(v))
	for i, x := range v {
		res[i] = uint16(x)
	}
	return
}
func SliceUint32(i any) (res []uint32) {
	v, _ := ToInt64SliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]uint32, len(v))
	for i, x := range v {
		res[i] = uint32(x)
	}
	return
}
func SliceUint(i any) (res []uint) {
	v, _ := ToInt64SliceE(i)
	if len(v) == 0 {
		return
	}
	res = make([]uint, len(v))
	for i, x := range v {
		res[i] = uint(x)
	}
	return
}

// ToSliceE casts an empty interface to a []any.
func ToSliceE(i any) ([]any, error) {

	var s []any

	switch v := i.(type) {
	case []any:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	case []map[string]any:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("Unable to Cast %#v of type %v to []any", i, reflect.TypeOf(i))
	}
}

// ToBoolSliceE casts an empty interface to a []bool.
func ToBoolSliceE(i any) ([]bool, error) {

	if i == nil {
		return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("Unable to Cast %#v to []bool", i)
	}
}

// ToStringSliceE casts an empty interface to a []string.
func ToStringSliceE(i any) ([]string, error) {

	var a []string

	switch v := i.(type) {
	case []any:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []string:
		return v, nil
	case string:
		return strings.Fields(v), nil
	case any:
		str, err := ToStringE(v)
		if err != nil {
			return a, fmt.Errorf("Unable to Cast %#v to []string", i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("Unable to Cast %#v to []string", i)
	}
}

// ToIntSliceE casts an empty interface to a []int.
func ToIntSliceE(i any) ([]int, error) {

	if i == nil {
		return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToIntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}
}

// ToIntSliceE casts an empty interface to a []int.
func ToInt64SliceE(i any) ([]int64, error) {

	if i == nil {
		return []int64{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToInt64E(s.Index(j).Interface())
			if err != nil {
				return []int64{}, fmt.Errorf("Unable to Cast %#v to []int", i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int64{}, fmt.Errorf("Unable to Cast %#v to []int", i)
	}
}
