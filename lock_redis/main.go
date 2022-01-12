package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	client     *redis.Client
	counterKey = "counter"
	lockKey    = "counter_locker"
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})
}

func increase() {
	// lock
	resp := client.SetNX(lockKey, 1, time.Second*2)

	lockKeySuccess, err := resp.Result()
	if err != nil || !lockKeySuccess {
		fmt.Printf("lock failed, lockKeySuccess: %v, err: %v\n", lockKeySuccess, err)
		return
	}

	getResp := client.Get(counterKey)
	cntValue, err := getResp.Int64()
	if err == nil || err == redis.Nil {
		cntValue++
		resp := client.Set(counterKey, cntValue, 0)
		_, err = resp.Result()
		if err != nil {
			fmt.Printf("set value failed, key: %v, value: %v, err: %v\n", counterKey, cntValue, err)
		}
	}

	fmt.Printf("current counter is %v\n", cntValue)

	delResp := client.Del(lockKey)
	unlockSuccess, err := delResp.Result()
	if err == nil && unlockSuccess > 0 {
		fmt.Printf("unlock success, lock key: %v\n", lockKey)
	} else if err != nil {
		fmt.Printf("unlock failed, err: %v\n", err)
		return
	}

}

func main() {
	// wait group
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increase()
		}()
	}

	wg.Wait()
}
