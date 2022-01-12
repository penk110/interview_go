package main

import (
	"fmt"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

func main() {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: "10.13.149.253:6379",
				Tube:     "tube",
			},
			{
				Endpoint: "10.13.149.253:6379",
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: "10.13.149.253:6379",
			Type: redis.NodeType,
		},
	})
	consumer.Consume(func(body []byte) {
		fmt.Println(string(body))
	})
}
