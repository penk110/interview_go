package main

import (
	"fmt"
	"time"
)

func put(m map[string]interface{}) {

	m["timeSp"] = time.Now().Unix()
	fmt.Printf("put ptr: %p, value: %v\n", m, m)

}

func main() {
	var bm = map[string]interface{}{
		"book":   "黑鹰传奇",
		"author": "戊戟先生",
	}

	fmt.Printf("before ptr: %p, value: %v\n", bm, bm)
	put(bm)
	fmt.Printf("after ptr: %p, value: %v\n", bm, bm)

	/*
		before ptr: 0xc000074180, value: map[author:戊戟先生 book:黑鹰传奇]
		put ptr: 0xc000074180, value: map[author:戊戟先生 book:黑鹰传奇 timeSp:1648805495]
		after ptr: 0xc000074180, value: map[author:戊戟先生 book:黑鹰传奇 timeSp:1648805495]
	*/
}
