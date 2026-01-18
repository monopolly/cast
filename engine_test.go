package cast

//testing
import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestConvert(u *testing.T) {
	__(u)

	type Test struct {
		Float float64
		Int   int
		Int8  int8
		Int16 int16
		Int32 int32
		Int64 int64

		Uint   uint
		Uint8  uint8
		Uint16 uint16
		Uint32 uint32
		Uint64 uint64

		MapStringAny   map[string]any
		MapStringBytes map[string][]byte
	}

	p := &Test{
		MapStringAny: map[string]any{},
	}

	var from any
	from = int(1)
	Convert(&p.Int, from)
	fmt.Println(p.Int, p.Int == 1)

	from = map[string]any{"k": "v", "k1": 1, "k2": true}
	Convert(&p.MapStringAny, from)
	fmt.Println(p.MapStringAny)

	from = map[string][]byte{"k": []byte("v")}
	Convert(&p.MapStringBytes, from)
	fmt.Println(p.MapStringBytes)
	select {}
}

func Benchmark1(u *testing.B) {
	u.ReportAllocs()

	for u.Loop() {

	}
}

func Benchmark2(u *testing.B) {
	u.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}

func __(u *testing.T) {
	fmt.Printf("\033[1;32m%s\033[0m\n", strings.ReplaceAll(u.Name(), "Test", ""))
}

func cmd(name string, v ...string) {
	c := exec.Command(name, v...)
	r, err := c.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
}
