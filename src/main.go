package main

import (
	"go_lb/domain/lb"
	"go_lb/domain/node"
)

func main() {
	node1 := node.NewNode("http://node1", 1024*1024, 100)
	node2 := node.NewNode("http://node2", 2*1024*1024, 150)

	loadBalancer := lb.NewLoadBalancer([]*node.Node{node1, node2})

	err := loadBalancer.HandleRequest(500)
	if err != nil {
		panic(err)
	}
}
