package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	count := math.MaxInt16
	for c := 0; c <= count; c++ {
		go func(i int64) {
			fmt.Printf("task<%v> ...\n", i)
			time.Sleep(time.Microsecond * 300)
		}(int64(c))
	}
}
