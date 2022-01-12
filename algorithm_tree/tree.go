package main

import "fmt"

/*
	树
*/

type node struct {
	ID        int
	Value     string
	LeftNode  *node
	RightNode *node
}

func (n *node) SetValue(v string) {
	n.Value = v
}

func (n *node) Traverse() {
	if n == nil {
		return
	}
	n.LeftNode.Traverse()
	fmt.Println()
	n.RightNode.Traverse()

}

func main() {

}
