package main

import (
	"bufio"
	"fmt"
	"github.com/Centny/gwf/log"
	"github.com/Centny/gwf/util"
	"io"
	oLog "log"
	"os"
)

const (
	Ldate         = 1 << iota // the date in the local time zone: 2009/01/23
	Ltime                     // the time in the local time zone: 01:23:23
	Lmicroseconds             // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                 // full file name and line number: /a/b/c/d.go:23
	Lshortfile                // final file name element and line number: d.go:23. overrides Llongfile

	D_LOG_FLAGS int = Ldate | Lmicroseconds | Lshortfile
)

func main() {
	//create and open log file
	path := "interview_go/io_log_stdout/tmp.log"
	util.FTouch(path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
		return
	}
	//create file buffer writer.
	//bo := bufio.NewWriter(f)
	bo := bufio.NewWriter(os.Stdout)

	defer f.Close()
	defer bo.Flush()
	//set the log writer to file and stdout.
	//or log.SetWriter(bo) to only file.
	log.SetWriter(io.MultiWriter(bo, os.Stdout))
	//set the log level
	log.SetLevel(log.DEBUG)
	oLog.SetOutput(os.Stdout)
	//use
	log.I("tesing")
	log.D("tesing")
	fmt.Println("print tesing")
	oLog.Println("origin log print tesing")

	/*
		1 4 16 或运算得到：21

		0000 0001
		0000 0100
		0001 0000
		或运算得到：
		0001 0101 = 1+4+16 = 21
	*/
	oLog.Println(Ldate, Lmicroseconds, Lshortfile, D_LOG_FLAGS)
}
