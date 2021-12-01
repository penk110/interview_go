package main

import "fmt"

func main() {
	s1 := make([]int, 10)
	s2 := make([]int, 0, 10)

	// [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(s1, len(s1), cap(s1))
	// []
	fmt.Println(s2, len(s2), cap(s2))
}
