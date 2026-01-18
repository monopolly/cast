// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"reflect"
)

func Map(i any, h func(k, v any)) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			h(key.Interface(), value.Interface())
		}
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func MapIntString(i any) map[int]string {
	var m = make(map[int]string)
	switch v := i.(type) {
	case map[int]string:
		return v
	default:
		Map(i, func(k, v any) {
			m[Int(k)] = fmt.Sprint(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func MapIntAny(i any) map[int]any {
	var m = make(map[int]any)
	switch v := i.(type) {
	case map[int]any:
		return v
	default:
		Map(i, func(k, v any) {
			m[Int(k)] = v
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func MapIntInt(i any) map[int]int {
	var m = make(map[int]int)
	switch v := i.(type) {
	case map[int]int:
		return v
	default:
		Map(i, func(k, v any) {
			m[Int(k)] = Int(v)
		})
		return m
	}
}

func MapIntBool(i any) map[int]bool {
	var m = make(map[int]bool)
	switch v := i.(type) {
	case map[int]bool:
		return v
	default:
		Map(i, func(k, v any) {
			m[Int(k)] = Bool(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func MapStringInt(i any) map[string]int {
	var m = make(map[string]int)
	switch v := i.(type) {
	case map[string]int:
		return v
	default:
		Map(i, func(k, v any) {
			m[String(k)] = Int(v)
		})
		return m
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func MapStringFloats(i any) map[string]float64 {
	var m = make(map[string]float64)
	switch v := i.(type) {
	case map[string]float64:
		return v
	default:
		Map(i, func(k, v any) {
			m[String(k)] = Float(v)
		})
		return m
	}
}

func MapStringString(i any) map[string]string {
	var m = make(map[string]string)
	switch v := i.(type) {
	case map[string]string:
		return v
	default:
		Map(i, func(k, v any) {
			m[String(k)] = String(v)
		})
		return m
	}
}

func StringMapStringSlice(i any) map[string][]string {
	v, _ := ToStringMapStringSliceE(i)
	return v
}

func StringMapBool(i any) map[string]bool {
	v, _ := ToStringMapBoolE(i)
	return v
}

func MapStringAny(i any) map[string]any {
	v, _ := ToStringMapE(i)
	return v
}

// ToStringMapStringE casts an empty interface to a map[string]string.
func ToStringMapStringE(i any) (map[string]string, error) {

	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]any:
		for k, val := range v {
			m[k] = String(val)
		}
		return m, nil
	case map[string]int:
		for k, val := range v {
			m[k] = fmt.Sprint(val)
		}
		return m, nil
	case map[string]int64:
		for k, val := range v {
			m[k] = fmt.Sprint(val)
		}
		return m, nil
	case map[int64]int64:
		for k, val := range v {
			m[fmt.Sprint(k)] = fmt.Sprint(val)
		}
		return m, nil
	case map[int]int64:
		for k, val := range v {
			m[fmt.Sprint(k)] = fmt.Sprint(val)
		}
		return m, nil
	case map[int64]string:
		for k, val := range v {
			m[fmt.Sprint(k)] = fmt.Sprint(val)
		}
		return m, nil
	case map[int64]any:
		for k, val := range v {
			m[fmt.Sprint(k)] = fmt.Sprint(val)
		}
		return m, nil
	case map[any]string:
		for k, val := range v {
			m[fmt.Sprint(k)] = val
		}
		return m, nil
	case map[any]any:
		for k, val := range v {
			m[fmt.Sprint(k)] = fmt.Sprint(val)
		}
		return m, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]string", i)
	}
}

// ToStringMapStringSliceE casts an empty interface to a map[string][]string.
func ToStringMapStringSliceE(i any) (map[string][]string, error) {

	var m = map[string][]string{}

	switch v := i.(type) {
	case map[string][]string:
		return v, nil
	case map[string][]any:
		for k, val := range v {
			m[String(k)] = SliceString(val)
		}
		return m, nil
	case map[string]string:
		for k, val := range v {
			m[String(k)] = []string{val}
		}
	case map[string]any:
		for k, val := range v {
			switch vt := val.(type) {
			case []any:
				m[String(k)] = SliceString(vt)
			case []string:
				m[String(k)] = vt
			default:
				m[String(k)] = []string{String(val)}
			}
		}
		return m, nil
	case map[any][]string:
		for k, val := range v {
			m[String(k)] = SliceString(val)
		}
		return m, nil
	case map[any]string:
		for k, val := range v {
			m[String(k)] = SliceString(val)
		}
		return m, nil
	case map[any][]any:
		for k, val := range v {
			m[String(k)] = SliceString(val)
		}
		return m, nil
	case map[any]any:
		for k, val := range v {
			key, err := ToStringE(k)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			value, err := ToStringSliceE(val)
			if err != nil {
				return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
			}
			m[key] = value
		}
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string][]string", i)
	}
	return m, nil
}

// ToStringMapBoolE casts an empty interface to a map[string]bool.
func ToStringMapBoolE(i any) (map[string]bool, error) {

	var m = map[string]bool{}

	switch v := i.(type) {
	case map[any]any:
		for k, val := range v {
			m[String(k)] = Bool(val)
		}
		return m, nil
	case map[string]any:
		for k, val := range v {
			m[String(k)] = Bool(val)
		}
		return m, nil
	case map[string]bool:
		return v, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]bool", i)
	}
}

// ToStringMapE casts an empty interface to a map[string]any.
func ToStringMapE(i any) (map[string]any, error) {

	var m = map[string]any{}

	switch v := i.(type) {
	case map[any]any:
		for k, val := range v {
			m[String(k)] = val
		}
		return m, nil
	case map[string]any:
		return v, nil
	default:
		return m, fmt.Errorf("Unable to Cast %#v to map[string]any", i)
	}
}

func MapStringBytes(i any) (res map[string][]byte) {
	res = map[string][]byte{}
	MapIterate(i, func(k, v any) { res[String(k)] = Bytes(v) })
	return
}

func MapStringBool(i any) (res map[string]bool) {
	res = map[string]bool{}
	MapIterate(i, func(k, v any) { res[String(k)] = Bool(v) })
	return
}

func IsMap(v any) bool {
	if v == nil {
		return false
	}
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func MapIterate(v any, f func(k, v any)) {
	// Проверяем, что это мапа
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return
	}

	// Проверяем, что мапа не nil
	if rv.IsNil() {
		return
	}

	// Итерируемся по элементам
	iter := rv.MapRange()
	for iter.Next() {
		f(iter.Key().Interface(), iter.Value().Interface())
	}

}
