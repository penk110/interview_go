package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

/*
	作为参数传递时，为引用传递
*/

func Worker(wid string, c chan string) {
	for v := range c {
		fmt.Printf("worker id: %s rcv: %v\n", wid, v)
	}
}

func Master(id string) chan string {
	var c = make(chan string)
	go Worker(id, c)

	return c
}

func main() {
	c := Master("123")
	ticker := time.NewTicker(time.Second)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			c <- strconv.FormatInt(time.Now().Unix(), 10)
		case <-ctx.Done():
			return
		}

	}
}
