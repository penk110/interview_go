package main

import "fmt"

func main() {
	var (
		ui uint32
		ei interface{}
	)

	ui = 32
	ei = ui

	switch ei.(type) {
	case uint16:
		ui = 8
	case uint32:
		ui = 3232
	}

	ui2 := ei.(uint32)
	fmt.Println(ui2)
}
