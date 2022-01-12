package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("线程安全的map\n")

	m := sync.Map{}
	m.Store("address", "127.0.0.1")
	m.Store("port", 8000)
	m.Store("db", "dev")
	db, ok := m.Load("db")
	if !ok {
		fmt.Println("get value failed")
	}
	fmt.Printf("sync map: %v\n", db)
	m.LoadOrStore("db", "dev123")
	m.Store("db", "dev")
}
