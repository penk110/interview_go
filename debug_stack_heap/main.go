package main

import (
	"fmt"
	"net"
	"strings"

	"runtime"
	"time"
)

func main() {
	fmt.Printf("[ResolveHostIp] ip: %v\n", ResolveHostIp())
	SetHost("key", "value")
}

// ResolveHostIp fech local ip
func ResolveHostIp() (ip string) {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ip
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip = networkIp.IP.String()
			return ip
		}
	}
	return ip
}

// SetHost set local ip to redis
func SetHost(k string, v string) (ok bool) {
	if v == ":80" {
		v = "http://" + ResolveHostIp() + v
	} else if !strings.HasPrefix(v, "http://") {
		v = "http://" + v
	}
	go func(key string, value interface{}) {
		ticker := time.NewTicker(time.Second * 60)
		for {
			select {
			case <-ticker.C:
				_ = SetHost(k, v)
			default:
				time.Sleep(time.Millisecond * 1000)
				runtime.Gosched()
			}
		}
	}(k, v)
	return true
}
