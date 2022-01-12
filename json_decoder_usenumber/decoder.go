package main

import (
	"encoding/json"
	"log"
	"strings"
)

func UseNumberDecoder(v string, data interface{}) (err error) {
	if v == "" {
		return nil
	}

	decoder := json.NewDecoder(strings.NewReader(v))
	decoder.UseNumber()
	return decoder.Decode(&data)
}

func main() {
	data := make(map[string]interface{})
	v := `{"f":"decoder","p":888927892776821}`

	err := UseNumberDecoder(v, &data)
	if err != nil {
		log.Fatalf("[UseNumberDecoder] err: %v\n", err.Error())
		return
	}

	log.Printf("[UseNumberDecoder] data: %v\n", data)

	data2 := make(map[string]interface{})
	_ = json.Unmarshal([]byte(v), &data2)
	log.Printf("data2: %v\n", data2)

	data3 := make(map[string]interface{})
	v3 := `{"f":"decoder3","p":888927892776821.1}`

	err = UseNumberDecoder(v3, &data3)
	if err != nil {
		log.Fatalf("[UseNumberDecoder] err: %v\n", err.Error())
		return
	}
	log.Printf("[UseNumberDecoder] data3: %v\n", data3)

	/*
		[UseNumberDecoder] data: map[f:decoder p:888927892776821]
		data2: map[f:decoder p:8.88927892776821e+14]
		[UseNumberDecoder] data3: map[f:decoder3 p:888927892776821.1]
	*/
}
