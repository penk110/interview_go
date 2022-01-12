package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/sony/sonyflake"
)

// getMachineID get machine id and return
func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}

	return machineID, nil
}

// readMachineIDFromLocalFile read machine id from local file and return
func readMachineIDFromLocalFile() uint16 {
	// logics
	return uint16(rand.Uint32())
}

// generateMachineID generate machine id and return
func generateMachineID() (uint16, error) {
	// logics
	return uint16(rand.Uint32()), nil
}

// checkMachineID check machine id and return bool value
func checkMachineID(machineID uint16) bool {
	// logics
	return true
}

/*
func checkMachineID(machineID uint16) bool {
	addResult, err := addMachineIDToRedisSet()
	if err != nil || addResult == 0 {
		return true
	}

	err := saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}

	return false

}

*/

func main() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")

	// create sonyFlake setting instance
	setting := sonyflake.Settings{
		StartTime:      t,
		MachineID:      generateMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(setting)
	//fmt.Println(sf.NextID())
	//for i := 1; i <= 10; i++ {
	//	id, err := sf.NextID()
	//	if err != nil {
	//		fmt.Printf("i: %v, err: %v\n", i, err)
	//		continue
	//	}
	//	// base 配置选择生成的进制类型 2-进制 16-进制 32-进制
	//	sid := strconv.FormatUint(id, 16)
	//	fmt.Printf("i: %v\tid: %v\tsid: %s\tsidLen: %d\n", i, id, strings.ToUpper(sid), len(sid))
	//
	//}

	// 奇妙的打印排序问题
	runtime.GOMAXPROCS(1)
	for i := 1; i <= 9; i++ {
		go func(i int) {
			time.Sleep(time.Microsecond * 800)
			id, err := sf.NextID()
			if err != nil {
				fmt.Printf("i: %v, err: %v\n", i, err)
			}
			// base 配置选择生成的进制类型 2-进制 16-进制 32-进制
			sid := strconv.FormatUint(id, 16)
			fmt.Printf("i: %v\tid: %v\tsid: %s\tsidLen: %d\n", i, id, strings.ToUpper(sid), len(sid))
			//time.Sleep(time.Microsecond * 800)
		}(i)
	}

	/*
		i: 9    id: 180047649548796994  sid: 27FA86222000442    sidLen: 15
		i: 1    id: 180047649565574210  sid: 27FA86223000442    sidLen: 15
		i: 2    id: 180047649565639746  sid: 27FA86223010442    sidLen: 15
		i: 3    id: 180047649565705282  sid: 27FA86223020442    sidLen: 15
		i: 4    id: 180047649565770818  sid: 27FA86223030442    sidLen: 15
		i: 5    id: 180047649565836354  sid: 27FA86223040442    sidLen: 15
		i: 6    id: 180047649565901890  sid: 27FA86223050442    sidLen: 15
		i: 7    id: 180047649565967426  sid: 27FA86223060442    sidLen: 15
		i: 8    id: 180047649566032962  sid: 27FA86223070442    sidLen: 15
	*/
	time.Sleep(time.Second * 3)
}
