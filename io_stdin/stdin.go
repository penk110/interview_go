package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stdIn := os.Stdin
	buffer := [8]byte{}
	c1 := make(chan os.Signal, 2)
	signal.Notify(c1, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-c1:
			fmt.Printf("[SINGLE] recive single")
			return
		default:
			n, err := stdIn.Read(buffer[:])
			if err == io.EOF {
				break
			}
			if n > 0 {
				fmt.Println(n, string(buffer[:n]))
			}
		}
	}

}
