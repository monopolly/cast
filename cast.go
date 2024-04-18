// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"reflect"
	"time"
)

func Bool(i interface{}) bool {
	v, _ := ToBoolE(i)
	return v
}

func Time(i interface{}) time.Time {
	v, _ := ToTimeE(i)
	return v
}

func Duration(i interface{}) time.Duration {
	v, _ := ToDurationE(i)
	return v
}

func Float(i interface{}) float64 {
	v, _ := ToFloat64E(i)
	return v
}

func Int64(i interface{}) int64 {
	v, _ := ToInt64E(i)
	return v
}

func Uint64(i interface{}) uint64 {
	v, _ := ToInt64E(i)
	return uint64(v)
}

func Uint32(i interface{}) uint32 {
	v, _ := ToInt64E(i)
	return uint32(v)
}

func Int(i interface{}) int {
	v, _ := ToIntE(i)
	return v
}

func String(i interface{}) string {
	switch s := i.(type) {
	case string:
		return s
	case []byte:
		return string(s)
	case nil:
		return ""
	default:
		return fmt.Sprint(s)
	}
}

// uint8
func Byte(i interface{}) byte {
	switch s := i.(type) {
	case byte:
		return s
	default:
		v := Uint32(i)
		if v > 255 {
			return 0
		}
		return uint8(v)
	}
}

func Bytes(i interface{}) []byte {
	switch s := i.(type) {
	case []byte:
		return s
	case string:
		return []byte(s)
	case nil:
		return nil
	default:
		return []byte(fmt.Sprint(s))
	}
}

func Map(i interface{}, h func(k, v interface{})) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			h(key.Interface(), value.Interface())
		}
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
func MapIntString(i interface{}) map[int]string {
	var m = make(map[int]string)
	switch v := i.(type) {
	case map[int]string:
		return v
	default:
		Map(i, func(k, v interface{}) {
			m[Int(k)] = fmt.Sprint(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
func MapIntInt(i interface{}) map[int]int {
	var m = make(map[int]int)
	switch v := i.(type) {
	case map[int]int:
		return v
	default:
		Map(i, func(k, v interface{}) {
			m[Int(k)] = Int(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
func MapStringInt(i interface{}) map[string]int {
	var m = make(map[string]int)
	switch v := i.(type) {
	case map[string]int:
		return v
	default:
		Map(i, func(k, v interface{}) {
			m[String(k)] = Int(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]interface{}.
func MapStringFloats(i interface{}) map[string]float64 {
	var m = make(map[string]float64)
	switch v := i.(type) {
	case map[string]float64:
		return v
	default:
		Map(i, func(k, v interface{}) {
			m[String(k)] = Float(v)
		})
		return m
	}
}

func StringMapString(i interface{}) map[string]string {
	var m = make(map[string]string)
	switch v := i.(type) {
	case map[string]string:
		return v
	default:
		Map(i, func(k, v interface{}) {
			m[String(k)] = String(v)
		})
		return m
	}
}

func StringMapStringSlice(i interface{}) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

func StringMapBool(i interface{}) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

func StringMap(i interface{}) map[string]interface{} {
	v, _ := ToStringMapE(i)
	return v
}

func SliceInterface(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

func SliceBool(i interface{}) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

func SliceString(i interface{}) []string {
	v, _ := ToStringSliceE(i)
	return v
}

func SliceInt(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

func SliceInt64(i interface{}) []int64 {
	v, _ := ToInt64SliceE(i)
	return v
}
