package unittest_benchmark

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type BlackCopy struct {
	f1, f2, f3, f4, f5, f6, f7, f8, f9 float64
	f11, fqw, f13                      float64
}
type BlackCopy1 struct {
	f1         float64
	s1, s2, s3 string
	i1         int
	t1, t2     time.Time
}

func (bc BlackCopy) foo1() {}

func (bc BlackCopy1) foo1() {}

func BenchmarkCopy(b *testing.B) {
	var a = BlackCopy{}
	fmt.Println(unsafe.Sizeof(a))
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e5; j++ {
			a.foo1()
		}
	}
	fmt.Sprintln(a)
}
func BenchmarkCopy1(b *testing.B) {
	var a = BlackCopy1{}
	fmt.Println(unsafe.Sizeof(a))
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1e5; j++ {
			a.foo1()
		}
	}

	fmt.Sprintln(a)
}
