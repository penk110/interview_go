package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 10,
			"msg":  "success",
			"data": map[string]interface{}{
				"apiName": "test",
				"openId":  1234,
				"dataId":  "abc",
			},
		})
	})

	ginpprof.Wrapper(router)

	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}

/*
浏览器器中访问：
	127.0.0.1：8000/debug/pprof/

/debug/pprof/

Types of profiles available:
Count	Profile
2	allocs
0	block
0	cmdline
4	goroutine
2	heap
0	mutex
0	profile
7	threadcreate
0	trace
full goroutine stack dump
Profile Descriptions:

allocs: A sampling of all past memory allocations
block: Stack traces that led to blocking on synchronization primitives
cmdline: The command line invocation of the current program
goroutine: Stack traces of all current goroutines
heap: A sampling of memory allocations of live objects. You can specify the gc GET parameter to run GC before taking the heap sample.
mutex: Stack traces of holders of contended mutexes
profile: CPU profile. You can specify the duration in the seconds GET parameter. After you get the profile file, use the go tool pprof command to investigate the profile.
threadcreate: Stack traces that led to the creation of new OS threads
trace: A trace of execution of the current program. You can specify the duration in the seconds GET parameter. After you get the trace file, use the go tool trace command to investigate the trace.


*/
