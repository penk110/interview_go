package main

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/net"
)

func main() {
	var (
		connStats []net.ConnectionStat
		ioStats   []net.IOCountersStat
		err       error
	)
	// 获取当前网络连接信息
	// 可填入tcp、udp、tcp4、udp4等等
	if connStats, err = net.Connections("all"); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("connStats: ", connStats)

	// 获取网络读写字节／包的个数
	if ioStats, err = net.IOCounters(false); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("ioStats:", ioStats)
}
