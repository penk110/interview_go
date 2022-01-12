package main

import (
	"os"
	"runtime/trace"
)

func main() {
	// Start enables tracing for the current program.
	// While tracing, the trace will be buffered and written to w.
	// Start returns an error if tracing is already enabled.
	err := trace.Start(os.Stderr)
	if err != nil {
		panic(err)
	}

	// Stop stops the current tracing, if any.
	// Stop only returns after all the writes for the trace have completed.
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "TRACE"
	}()

	<-ch

}
