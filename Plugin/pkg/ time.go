package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("[bad pluin] init------")
}

// T struct t
type T struct {
}

// FmtPrint return fmt timeStr
func (t T) FmtPrint(tt time.Time) {
	fmt.Println("[FmtPrint] start------")
	fmt.Println(tt)
	fmt.Println("[FmtPrint] end------")
}

// OBJ global object
var OBJ = T{}
