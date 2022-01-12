package main

import (
	"fmt"
	"runtime"
)

type err struct {
}

func (e *err) Printf() {
	fmt.Printf("Printf func \n")
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("ok is false\n")
	}
	fmt.Printf("%v %v %v\n", pc, file, line)
}

func (e *err) Println() {
	fmt.Printf("Println func \n")
	// pc uintptr, file string, line int, ok bool
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("ok is false\n")
	}
	fmt.Printf("%v %v %v\n", pc, file, line)
}

func (e *err) Error() {
	fmt.Printf("Error func \n")
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("ok is false\n")
	}
	fmt.Printf("%v %v %v\n", pc, file, line)
}

func main() {
	e := new(err)
	e.Printf()
}
