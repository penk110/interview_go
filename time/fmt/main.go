package main

import (
	"fmt"
	"reflect"
)


func main() {
	s := "12"
	fmt.Printf("%v\n", reflect.TypeOf(s))
	i := 0
	fmt.Printf("%v\n", reflect.TypeOf(i))
}