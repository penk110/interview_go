package main

import (
	"fmt"
	"github.com/tal-tech/go-queue/dq"
	"strconv"
	"time"
)

func main() {
	producer := dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: "localhost:63791",
			Tube:     "tube",
		},
		{
			Endpoint: "localhost:63792",
			Tube:     "tube",
		},
	})

	for i := 1000; i < 1005; i++ {
		_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second*5)
		if err != nil {
			fmt.Println(err)
		}
	}
}
