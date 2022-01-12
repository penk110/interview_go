package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func init() {
	pwd, _ := os.Getwd()
	_ = os.Setenv("path", pwd)
}

func main() {
	s := []int{
		11, 22, 33, 44, 55, 55, 55, 55, 55, 55,
	}
	arrays := rand.Perm(len(s))
	fmt.Printf("perm: %v\n", arrays)
	path := os.Getenv("path")
	log.Printf("env: %v\n", path)
}
