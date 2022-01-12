package lib_psutils

import (
	"context"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	// "github.com/shirou/gopsutil/mem"  // to use v2
)

func TestVirtualMemory(t *testing.T) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	t.Logf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	t.Log(v)
}

func TestCPU(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()
	stat, err := cpu.InfoWithContext(ctx)
	if err != nil {
		t.Errorf("featch cpu stat failed, err: %v\n", err.Error())
		return
	}
	t.Log(stat)
}
