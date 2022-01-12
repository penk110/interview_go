package main

import (
	"encoding/json"
	"io"
)

/*
	检查变量是否为 nil 在进行操作

*/

type RequestOpts struct {
	Scheme    string            `json:"scheme"`
	URL       string            `json:"url"`
	Method    string            `json:"method"`
	Header    map[string]string `json:"header"` // Content-Type Cookie
	Query     map[string]string `json:"query"`
	Body      interface{}       `json:"body"`
	Proxy     string            `json:"proxy"`
	HTTP2     bool              `json:"http2"`
	Keepalive bool              `json:"keepalive"`
}

type RequestKVBody struct {
}

// Unmarshal 序列化前检查接收者是否为 空指针
func Unmarshal(buf []byte, receiver interface{}) error {
	if receiver == nil {
		return io.EOF
	}

	return json.Unmarshal(buf, receiver)
}

func main() {

}
