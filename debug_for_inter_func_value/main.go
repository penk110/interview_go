package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func() {
			// 闭包直接引用for循环中定义的i，最后一次循环i++后等于5不进入循环体内
			// 但是之前的goroutine内部是对i的引用，所以输出的值可能是最后一次自增得到的值
			fmt.Printf("task<%d> running...\n", i)
		}()
	}

	msg := struct {
		ID   int    `json:"id"`
		Desc string `json:"desc"`
	}{}
	m := map[int]interface{}{}
	mk := []int{}
	for i := 1; i <= 100; i++ {
		msg.ID = i
		msg.Desc = "desc_" + strconv.FormatInt(int64(i), 10)
		s, _ := json.Marshal(msg)
		if _, ok := m[i]; ok {
			mk = append(mk, i)
		} else {
			m[i] = nil
		}
		fmt.Println(i+1, " % 2 = ", (i+1)%2)
		fmt.Printf("%v\n", string(s))
	}
	fmt.Println(mk)

	time.Sleep(time.Second * 3)
}
