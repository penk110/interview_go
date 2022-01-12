package main

import (
	"fmt"
)

// 闭包引用验证

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := app()
	b := app()

	fmt.Println(a("go"))

	fmt.Println(b("All"))

}
