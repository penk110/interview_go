package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

//
const (
	mb = 1024 * 1024
	gb = 1024 * mb
)

func main() {
	// wait for all goroutine to finish
	wg := sync.WaitGroup{}

	channel := make(chan string)
	dist := make(map[string]int64)
	done := make(chan bool, 1)

	go func() {
		for s := range channel {
			dist[s] ++
		}
		done <- true
	}()

	var current int64
	//var limit int64 = 500 * gb
	var limit int64 = 2048
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int, c, l int64) {
			read(current, limit, "./main.go", channel)
			fmt.Printf("%d thread has been completed\n", i)
			wg.Done()
		}(i, current, limit)
		current += limit + 1
		fmt.Printf("current: %v, limit: %v\n", current, limit)
	}

	wg.Wait()
	close(channel)
	<-done
	close(done)
}

func readFile(offset int64, limit int64, fileName string, channel chan (string)) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	// Move the pointer of the file to the start of designated chunk.
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)

	// This block of code ensures that the start of chunk is a new word. If
	// a character is encountered at the given position it moves a few bytes till
	// the end of the word.
	if offset != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

		if err != nil {
			panic(err)
		}
	}
}

// read
func read(offset, limit int64, fileName string, channel chan string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file failed, err: ", err)
		return
	}
	defer file.Close()
	// move the pointer of the file to start of designated chunk
	ret, err := file.Seek(offset, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret)
	reader := bufio.NewReader(file)
	file.Seek(offset, 0)
	if offset != 0 {
		_, err = reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

		if err != nil {
			panic(err)
		}
	}

	var cumulativeSize int64
	for {
		// break
		if cumulativeSize > limit {
			break
		}

		b, err := reader.ReadBytes(' ')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		cumulativeSize += int64(len(b))
		s := strings.TrimSpace(string(b))
		if s != "" {
			channel <- s
		}

	}
}
