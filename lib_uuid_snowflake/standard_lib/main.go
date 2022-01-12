package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
	"time"
)

func main() {
	//nodeID := rand.Int63()
	sf, err := snowflake.NewNode(1)
	if err != nil {
		panic(sf)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := sf.Generate()
		fmt.Printf("id: %v\n", id)
		fmt.Printf("node: %v\nstep: %v\ntime:%v\n", id.Node(), id.Step(), id.Time())
		fmt.Printf("**********\n")
	}

	// 1288834974657
	t := time.Unix(1288834974657/1000, 0)
	fmt.Printf("time: %v\n", t.Format("2006-01-02 15:04:05"))
}

/*

var (
	// Epoch is set to the twitter snowflake epoch of Nov 04 2010 01:42:54 UTC in milliseconds
	// You may customize this to set a different epoch for your application.
	Epoch int64 = 1288834974657

	// NodeBits holds the number of bits to use for Node
	// Remember, you have a total 22 bits to share between Node/Step
	NodeBits uint8 = 10

	// StepBits holds the number of bits to use for Step
	// Remember, you have a total 22 bits to share between Node/Step
	StepBits uint8 = 12

	// DEPRECATED: the below four variables will be removed in a future release.
	mu        sync.Mutex
	nodeMax   int64 = -1 ^ (-1 << NodeBits)
	nodeMask        = nodeMax << StepBits
	stepMask  int64 = -1 ^ (-1 << StepBits)
	timeShift       = NodeBits + StepBits
	nodeShift       = StepBits
)

*/
