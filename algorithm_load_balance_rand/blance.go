package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ServerInstance struct {
	Name   string
	Addr   string
	Port   int
	Weight uint32
}

type ServerPool struct {
	instances []*ServerInstance
	mux       *sync.Mutex
	size      uint32
}

var serverPool *ServerPool

func init() {
	serverPool = &ServerPool{
		instances: make([]*ServerInstance, 0),
		mux:       &sync.Mutex{},
		size:      0,
	}
}
func (s *ServerPool) Register(ins ...*ServerInstance) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.instances = append(s.instances, ins...)
	s.size += uint32(len(ins))

}

func (s *ServerPool) Next() *ServerInstance {
	s.mux.Lock()
	defer s.mux.Unlock()

	// TODO: 随机取余
	instance := s.instances[rand.Uint32()%s.size]
	return instance
}

func main() {
	ins := []*ServerInstance{
		{
			Name:   "s1",
			Addr:   "192.168.1.10",
			Port:   8080,
			Weight: 1,
		},
		{
			Name:   "s2",
			Addr:   "192.168.1.10",
			Port:   8081,
			Weight: 2,
		},
		{
			Name:   "s3",
			Addr:   "192.168.1.10",
			Port:   8082,
			Weight: 3,
		},
	}

	serverPool.Register(ins...)

	for i := 0; i < 10; i++ {
		ins1 := serverPool.Next()
		fmt.Println(ins1.Name)
	}

	//ins1 := serverPool.Next()
	//fmt.Println(ins1.Name)
	//ins2 := serverPool.Next()
	//fmt.Println(ins2.Name)
}
