package main

import (
	"fmt"
	"os"
	"reflect"
)

func OpenFile() error {
	var err *os.PathError

	return err
}

func main() {
	err := OpenFile()
	// <nil> false *fs.PathError
	fmt.Println(err, err == nil, reflect.TypeOf(err))
}
