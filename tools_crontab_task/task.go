package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Task struct{}

func (t Task) Run() {
	fmt.Println("cron task")
}

func main() {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	c.AddJob("* * * * *", Task{})

	c.Start()

	select {}
}
