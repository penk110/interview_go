package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		ch       chan int
		signalCh chan struct{}
		exit     chan os.Signal
		signalF  os.Signal
		i        int
		j        int
	)

	ch = make(chan int, 2)
	signalCh = make(chan struct{}, 1)
	exit = make(chan os.Signal)
	go func() {
	loop:
		for {
			select {
			case i = <-ch:
				log.Printf("recive: %v\n", i)
			case <-signalCh:
				println("channel has been closed!")
				break loop
			}
		}
	}()

	go func() {
		for {
			ch <- j
			j++
			time.Sleep(time.Second * 2)
			if j > 10 {
				signalCh <- struct{}{}
				return
			}
		}
	}()

	time.AfterFunc(time.Second*25, func() {
		exit <- os.Kill
	})
	signal.Notify(exit, os.Kill, os.Interrupt)
	signalF = <-exit
	log.Printf("os signal: %v\n", signalF.String())
}
