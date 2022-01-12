package main

import (
	"fmt"
	"log"
	"os"
	"plugin"
	"time"
)

// NormalT normal t struct
type NormalT interface {
	FmtPrint(tt time.Time)
}

func init() {

}

func main() {
	log.Println("main function stared")
	// load module 插件您也可以使用go http.Request从远程下载到本地,在加载做到动态的执行不同的功能
	// 1. open the so file to load the symbols
	plug, err := plugin.Open("pkg/time.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	log.Println("plugin opened")

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	doc, err := plug.Lookup("Doctor")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type GoodDoctor (defined above)
	t, ok := doc.(NormalT)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(2)
	}

	// 4. use the module
	t.FmtPrint(time.Now())
}
