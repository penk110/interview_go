package main

type Shape interface {
	Area() float64
}

type Rectangle struct {
	a float64
	b float64
}

// Area 标记禁止函数内联
//go:noinline
func (rectangle *Rectangle) Area() float64 {
	return rectangle.a * rectangle.b
}

func main() {
	var (
		r  *Rectangle
		ri Shape
	)
	r = &Rectangle{
		a: 10,
		b: 20,
	}

	ri = Shape(r)

	ri.Area()
}
