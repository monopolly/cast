// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"reflect"
)

// "to" must be pointer to type
func Convert(to, from any) (err error) {
	if to == nil {
		return fmt.Errorf("to is nil")
	}

	toinderected := indirect(to)

	// type
	switch toinderected.(type) {

	// int
	case int:
		set(to, Int(from))
	case int8:
		set(to, Int8(from))
	case int16:
		set(to, Int16(from))
	case int32:
		set(to, Int32(from))
	case int64:
		set(to, Int64(from))
	case uint:
		set(to, Uint(from))
	case uint8:
		set(to, Uint8(from))
	case uint16:
		set(to, Uint16(from))
	case uint32:
		set(to, Uint32(from))
	case uint64:
		set(to, Uint64(from))
	case float32:
		set(to, Float32(from))
	case float64:
		set(to, Float64(from))
	case string:
		set(to, String(from))
	case []byte:
		set(to, Bytes(from))
	case bool:
		set(to, Bool(from))

	// // maps int
	case map[string]any:
		set(to, MapStringAny(from))
	case map[string]string:
		set(to, MapStringString(from))
	case map[string]int:
		set(to, MapStringInt(from))
	case map[string]bool:
		set(to, MapStringBool(from))
	case map[string][]byte:
		set(to, MapStringBytes(from))

	// // map int int
	case map[int]int:
		set(to, MapIntInt(from))
	case map[int]bool:
		set(to, MapIntBool(from))
	case map[int]string:
		set(to, MapIntString(from))
	case map[int]any:
		set(to, MapIntAny(from))

	case []int:
		set(to, SliceInt(from))
	case []string:
		set(to, SliceString(from))
	case []any:
		set(to, SliceAny(from))

	case any:
		set(to, from)

	default:
		return fmt.Errorf("type is not supported")
	}

	return
}

// check support type name
func Supported(types string) (err error) {
	if types == "" {
		return fmt.Errorf("fail")
	}

	// type
	switch types {
	case "int":
	case "int8":
	case "int16":
	case "int32":
	case "int64":
	case "uint":
	case "uint8":
	case "uint16":
	case "uint32":
	case "uint64":
	case "float32":
	case "float64":
	case "string":
	case "[]byte":
	case "bool":
	case "map[string]any":
	case "map[string]string":
	case "map[string]int":
	case "map[string]bool":
	case "map[string][]byte":
	case "map[int]int":
	case "map[int]bool":
	case "map[int]string":
	case "map[int]any":
	case "[]int":
	case "[]string":
	case "[]any":
	case "any":
	default:
		return fmt.Errorf("type is not supported")
	}

	return
}

// check type support
func Support(to any) (err error) {
	if to == nil {
		return fmt.Errorf("to is nil")
	}

	toinderected := indirect(to)

	// type
	switch toinderected.(type) {

	// int
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
	case string:
	case []byte:
	case bool:
	case map[string]any:
	case map[string]string:
	case map[string]int:
	case map[string]bool:
	case map[string][]byte:
	case map[int]int:
	case map[int]bool:
	case map[int]string:
	case map[int]any:
	case []int:
	case []string:
	case []any:
	case any:
	default:
		return fmt.Errorf("type is not supported")
	}

	return
}

func set(to, from any) error {
	// Получаем reflect.Value для обеих переменных
	toVal := reflect.ValueOf(to)
	fromVal := reflect.ValueOf(from)

	// Проверяем, что 'to' - это указатель
	if toVal.Kind() != reflect.Pointer {
		return fmt.Errorf("'to' должен быть указателем, а не %v", toVal.Kind())
	}

	// Проверяем, что указатель не nil
	if toVal.IsNil() {
		return fmt.Errorf("'to' не может быть nil указателем")
	}

	// Получаем значение, на которое указывает 'to'
	elem := toVal.Elem()

	// Проверяем совместимость типов
	if !fromVal.Type().AssignableTo(elem.Type()) {
		// Попробуем преобразовать, если это возможно
		if fromVal.Type().ConvertibleTo(elem.Type()) {
			elem.Set(fromVal.Convert(elem.Type()))
			return nil
		}
		return fmt.Errorf("нельзя присвоить %v в %v",
			fromVal.Type(), elem.Type())
	}

	// Присваиваем значение
	elem.Set(fromVal)
	return nil
}
