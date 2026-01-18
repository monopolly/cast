// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
	"strconv"
)

func Bool(i any) bool {
	v, _ := ToBoolE(i)
	return v
}

// ToBoolE casts an empty interface to a bool.
func ToBoolE(i any) (bool, error) {

	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	case []uint8:
		return strconv.ParseBool(string(b))
		//return strconv.ParseBool(string(i.([]uint8)))
	default:
		return false, fmt.Errorf("Unable to Cast %#v to bool", i)
	}
}
