// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package cast

import (
	"fmt"
)

// uint8
func Byte(i any) byte {
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

func Bytes(i any) []byte {
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
