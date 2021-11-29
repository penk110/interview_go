package main

import "fmt"

func set1(a *string) {
	b := "set_b"
	fmt.Printf("1 %p %p\n", a, &b)
	a = &b
	fmt.Printf("2 %p %p %v %v\n", a, &b, *a, b)
	*a = b
	fmt.Printf("3 %p %p %v %v\n", a, &b, *a, b)
}

func set2(a **string) {
	b := "set_b"
	fmt.Printf("1 %p %p\n", a, &b)
	*a = &b
	fmt.Printf("2 %p %p %v %v\n", a, &b, *a, b)
	**a = b
	fmt.Printf("3 %p %p %v %v\n", a, &b, *a, b)
}

func main() {
	var a *string
	fmt.Println(a, &a)
	set1(a)
	fmt.Println(a, &a)

	// 指针的指针？？？？
	set2(&a)
	fmt.Println(a, &a, *a)
}
