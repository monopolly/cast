package cast

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount_Marshal(ggggg *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fn = fn[strings.LastIndex(fn, ".Test")+5:]
	fn = strings.Join(strings.Split(fn, "_"), ": ")
	fmt.Printf("\033[1;32m%s\033[0m\n", fn)

	a := assert.New(ggggg)
	_ = a

	b := Bool([]byte("t"))
	a.True(b)

	i64 := Int64([]byte("17971391732"))
	a.Equal(int64(17971391732), i64)

	i64 = Int64("17971391732")
	a.Equal(int64(17971391732), i64)

	i64 = Int64(17971391732)
	a.Equal(int64(17971391732), i64)
}

func BenchmarkNew(bbbbbbbb *testing.B) {
	bbbbbbbb.ReportAllocs()
	bbbbbbbb.ResetTimer()
	for n := 0; n < bbbbbbbb.N; n++ {

	}
}

func BenchmarkGetFreeParallel(bbbbbbbb *testing.B) {
	bbbbbbbb.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}
