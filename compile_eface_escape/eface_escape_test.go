package main

import "testing"

func BenchmarkUint32(b *testing.B) {
	var u uint32
	b.Run("uint32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			u = uint32(i)
		}
	})
	b.Logf("u: %v", u)

	var ueface interface{}
	b.Run("eface32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ueface = uint32(i)
		}
	})
	b.Logf("ueface: %v", ueface)
}

func BenchmarkUint16(b *testing.B) {
	var u uint16
	b.Run("uint16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			u = uint16(i)
		}
	})
	b.Logf("u: %v", u)

	var ueface interface{}
	b.Run("uint16", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ueface = uint16(i)
		}
	})
	b.Logf("ueface: %v", ueface)
}

func BenchmarkUint8(b *testing.B) {
	var u uint8
	b.Run("uint8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			u = uint8(i)
		}
	})
	b.Logf("u: %v", u)

	var ueface interface{}
	b.Run("uint8", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ueface = uint8(i)
		}
	})
	b.Logf("ueface: %v", ueface)
}

/*
空接口转换接口性能对比

➜  compile_eface_escape git:(master) ✗ go test -bench=.  -benchmem
goos: darwin
goarch: amd64
pkg: github.com/penk110/interview_go/compile_eface_escape
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkUint32/uint32-12               1000000000               0.2542 ns/op          0 B/op          0 allocs/op
BenchmarkUint32/eface32-12              98167009                11.21 ns/op            4 B/op          0 allocs/op
BenchmarkUint16/uint16-12               1000000000               0.2550 ns/op          0 B/op          0 allocs/op
BenchmarkUint16/uint16#01-12            100000000               11.10 ns/op            1 B/op          0 allocs/op
BenchmarkUint8/uint8-12                 1000000000               0.2552 ns/op          0 B/op          0 allocs/op
BenchmarkUint8/uint8#01-12              1000000000               0.7579 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/penk110/interview_go/compile_eface_escape    4.191s
*/
