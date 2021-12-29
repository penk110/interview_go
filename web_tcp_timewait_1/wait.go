package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func startReq() {
	count := 0
	for {
		resp, err := http.Get("http://192.168.1.18:8030/")
		if err != nil {
			panic(fmt.Sprintf("Got error: %v", err))
		}
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		log.Printf("Finished GET request #%v", count)
		count += 1
	}

}

func main() {
	for i := 0; i < 100; i++ {
		go startReq()
	}

	select {}
}
