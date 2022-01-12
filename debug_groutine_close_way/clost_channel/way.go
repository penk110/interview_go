package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		ch      chan int
		exit    chan os.Signal
		signalF os.Signal
		i       int
		j       int
	)

	ch = make(chan int, 2)
	exit = make(chan os.Signal)
	go func() {
		for i = range ch {
			log.Printf("recive: %v\n", i)
		}
		defer func() { println("channel has been closed!") }()
	}()

	go func() {
		for {
			ch <- j
			j++
			time.Sleep(time.Second * 2)
			if j > 10 {
				close(ch)
				return
			}
		}
	}()

	signal.Notify(exit, os.Kill, os.Interrupt)
	signalF = <-exit
	log.Printf("os signal: %v\n", signalF.String())
}
