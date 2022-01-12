package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var (
		ch         chan int
		exit       chan os.Signal
		ctx        context.Context
		cancelFunc context.CancelFunc
		signalF    os.Signal
		i          int
		j          int
	)

	ch = make(chan int, 2)
	exit = make(chan os.Signal)
	ctx, cancelFunc = context.WithTimeout(context.TODO(), time.Second*10)
	defer cancelFunc()
	go func(ctx context.Context) {
	loop:
		for {
			select {
			case i = <-ch:
				log.Printf("recive: %v\n", i)
			case <-ctx.Done():
				println("channel has been closed!")
				break loop
			}
		}
	}(ctx)

	go func() {
		for {
			ch <- j
			j++
			time.Sleep(time.Second * 2)
			//if j > 5 {
			//	// TODO: 一个goroutine panic，所有goroutine遭殃
			//	panic("timeout ...")
			//	return
			//}
			if j > 10 {
				return
			}
		}
	}()

	time.AfterFunc(time.Second*15, func() {
		exit <- os.Kill
	})
	signal.Notify(exit, os.Kill, os.Interrupt)
	signalF = <-exit
	log.Printf("os signal: %v\n", signalF.String())
}
