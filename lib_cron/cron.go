package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

var _cron *cron.Cron

var _flag *F

type F struct {
	Mu    sync.Mutex
	Count int64
}

func (f *F) Add() {
	f.Mu.Lock()
	defer f.Mu.Unlock()
	f.Count++
}

func init() {
	_cron = cron.New()

	_flag = &F{
		Mu:    sync.Mutex{},
		Count: 0,
	}
}

func main() {
	var (
		entryID cron.EntryID
		err     error
	)
	entryID, err = _cron.AddFunc("@every 1s", func() {
		_flag.Mu.Lock()
		defer _flag.Mu.Unlock()
		_flag.Count++
		fmt.Printf("tick every 1 second, count: %d\n", _flag.Count)
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("entryID: %d\n", entryID)

	_cron.Start()

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		select {
		case <-ticker.C:
			break
		}
		_cron.AddFunc("@every 2s", func() {
			_flag.Mu.Lock()
			defer _flag.Mu.Unlock()
			_flag.Count++
			fmt.Printf("==== tick every 1 second, count: %d\n", _flag.Count)
		})
		_cron.Remove(entryID)

	}()

	run()
}

func run() {
	for {
		select {}
	}
}
