package basic_define_interface

type Rectangle struct {
	a float64
	b float64
}

func (rectangle *Rectangle) perimeter() float64 {
	return (rectangle.a + rectangle.b) * 2
}

func (rectangle *Rectangle) area() float64 {
	return rectangle.a * rectangle.b
}

func (rectangle *Rectangle) SetA(a float64) {
	rectangle.a = a
}

func (rectangle *Rectangle) GetA() float64 {
	return rectangle.a
}
