package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	dat := []byte(`{"name":"Go","age":10}`)
	result := make(map[string]interface{})
	err := json.Unmarshal(dat, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result["name"], result["age"])
}
