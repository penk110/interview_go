package basic_define_interface

import "testing"

func TestAngle(t *testing.T) {
	var (
		shape1 Shape // 接口定义 为接口的静态类型
		shape2 Shape
	)

	shape1 = &Rectangle{
		a: 10,
		b: 20,
	}
	// 接口的具体实现 为接口的动态类型
	// 接口动态调用表现出不同动态类型的行为
	t.Logf("[rectangle] perimeter: %.2f, area: %.2f\n", shape1.perimeter(), shape1.area())

	shape2 = &Triangle{
		a: 3,
		b: 4,
		c: 5,
	}
	//
	t.Logf("[triangle] perimeter: %.2f, area: %.2f\n", shape2.perimeter(), shape2.area())

	// 接口签名外的方法
	// panic info: shape1.SetA undefined (type Shape has no field or method SetA)
	// shape1.SetA(3)

	// i.(Type) Type 为具体实现的动态类型接口，通过此方式 获取存储在接口中动态类型的结构体
	t.Logf("[rectangle] before a: %.2f\n", shape1.(*Rectangle).GetA())
	shape1.(*Rectangle).SetA(50)
	t.Logf("[rectangle] after a: %.2f\n", shape1.(*Rectangle).GetA())

	t.Logf("[triangle] get c: %.2f\n", shape2.(*Triangle).GetC())

	var (
		shape3 *Rectangle
		ok     bool
	)

	shape3, ok = shape1.(*Rectangle)
	// [rectangle] shape1 point: 0xc0000a01a0, shape3 point: 0xc0000a01a0, ok: true
	t.Logf("[rectangle] shape1 point: %p, shape3 point: %p, ok: %t\n", shape1, shape3, ok)

	shape3, ok = shape2.(*Rectangle)
	// [triangle] shape2 point: 0xc000018090, shape3 point: 0x0, ok: false
	t.Logf("[triangle] shape2 point: %p, shape3 point: %p, ok: %t\n", shape2, shape3, ok)

}

/*
接口编译时优化逻辑：


func implements(t, iface *types.Type, m, samename **types.Field, ptr *int) bool {
	...
}

最坏时间复杂度 o(m+n)
*/
