package main

import "fmt"
import "reflect"

type A struct {
	Foo  string
	Foo2 B
}

type B struct {
	Name string
	Age  int
}

func (a *A) PrintFoo() {
	fmt.Println("Foo value is " + a.Foo)
}

type T struct {
	A1 int
	B2 string
	A  A
}

func main() {
	//a := A{
	//	Foo: "foo",
	//	Foo2: B{
	//		Name: "123",
	//		Age:  222,
	//	},
	//}
	//val := reflect.Indirect(reflect.ValueOf(a))
	//
	//fmt.Println(val.Type())
	//fmt.Println(val.Type().Field(0).Name)

	//v := reflect.ValueOf(a)
	//t := v.Type()
	//for i := 0; i < t.NumField(); i++ {
	//	name := t.Field(i).Name
	//	value := v.FieldByName(name).Interface()
	//	if v.Type().Kind() == reflect.Struct {
	//		name := t.Field(i).Name
	//		value := v.FieldByName(name).Interface()
	//		fmt.Println("-----", name, value)
	//
	//	}
	//
	//	fmt.Println(name, value)
	//}

	t := T{
		A1: 23,
		B2: "skidoo",
		A: A{
			Foo:  "00",
			Foo2: B{
				Name: "11",
				Age:  22,
			},
		},
	}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Type().Kind() == reflect.Struct {

		}
		fmt.Println(typeOfT.Field(i).Tag)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}


