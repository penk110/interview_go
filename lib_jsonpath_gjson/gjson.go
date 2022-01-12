package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var jsonData = `
{
	"data": [
		{
			"name": "golang",
			"type": "static"
		},
		{
			"name": "python",
			"type": "dynamic"
		}
	]
}
`

func main() {
	host := gjson.Get(jsonData, "data.1")
	fmt.Println(host.String())

	queryKey0 := gjson.Get(jsonData, "data.0.name")
	fmt.Println(queryKey0)
}
