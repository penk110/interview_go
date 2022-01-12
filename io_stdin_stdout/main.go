package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 标准输入拷贝到标准输出
	i, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(i)
}
