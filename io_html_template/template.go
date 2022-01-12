package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type book struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

func template1() {
	b := book{
		Name:     "黑影传奇",
		Category: "武侠小说",
	}
	t, err := template.ParseFiles("/Users/zhang/gopath/src/zyphub/interview_go/io_html_template/template.html")
	if err != nil {
		panic(err)
	}

	// Execute 内部是在原来原有 buf 的基础上进行追加
	// buf := []byte{0, 0, 0} => buf = append(buf, 1, 2) => []byte{0, 0, 0, 1, 2}
	bf := bytes.NewBuffer([]byte{})
	err = t.Execute(bf, b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[writer] data: %v\n", bf)
	fmt.Printf("[writer] data: %v\n", bf.String())
}

func main() {
	template1()

}
