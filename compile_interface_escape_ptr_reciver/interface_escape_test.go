package main

import "testing"

func BenchmarkRectangle_Area(b *testing.B) {
	r := Rectangle{
		a: 5,
		b: 10,
	}

	for i := 0; i < b.N; i++ {
		// TODO: 直接调用
		r.Area()
	}
}

func BenchmarkRectangle_Area2(b *testing.B) {
	r := &Rectangle{
		a: 5,
		b: 10,
	}

	for i := 0; i < b.N; i++ {
		// TODO: 接口调用
		Shape(r).Area()
	}
}

/*



 */
