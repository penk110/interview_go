package basic_define_interface

import "math"

type Triangle struct {
	a float64
	b float64
	c float64
}

func (triangle *Triangle) perimeter() float64 {
	return triangle.a + triangle.b + triangle.c
}

func (triangle *Triangle) area() float64 {
	var p float64
	p = triangle.perimeter() / 2
	return math.Sqrt(p * (p - triangle.a) * (p - triangle.b) * (p - triangle.c))
}

func (triangle *Triangle) GetC() float64 {

	return triangle.c
}
