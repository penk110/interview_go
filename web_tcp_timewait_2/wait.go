package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func startWebserver() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(time.Millisecond * 50)

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":8080", nil)

}

func startLoadTest() {
	count := 0
	for {
		resp, err := http.Get("http://localhost:8080/")
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

	// start a webserver in a goroutine
	startWebserver()

	for i := 0; i < 100; i++ {
		go startLoadTest()
	}

	time.Sleep(time.Second * 2400)

}
