package main

import (
	"fmt"
	"math"
	"os"
)

func nullPoint() error {
	var err *os.PathError = nil
	return err
}

func main() {
	var a *int = nil
	var b interface{} = nil
	// true true false
	fmt.Println(a == nil, b == nil, a == b)

	var i = math.MinInt32
	var j = uint32(i)
	//-2147483648 2147483648
	fmt.Println(i, j)

	p := nullPoint()
	//<nil> false
	fmt.Println(p, p == nil)

	// 类型
	fmt.Printf("%v %T", p, p)

}
