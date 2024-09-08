package lb

import (
	"fmt"
	"go_lb/domain/node"
)

type LoadBalancer struct {
	Nodes        []*node.Node
	CurrentIndex int
}

func NewLoadBalancer(nodes []*node.Node) *LoadBalancer {
	return &LoadBalancer{
		Nodes:        nodes,
		CurrentIndex: 0,
	}
}

func (lb *LoadBalancer) SelectNode(requestSize int) (*node.Node, error) {
	startIndex := lb.CurrentIndex
	for {
		node := lb.Nodes[lb.CurrentIndex]
		if node.CanHandleRequest(requestSize) {
			lb.CurrentIndex = (lb.CurrentIndex + 1) % len(lb.Nodes)
			return node, nil
		}

		lb.CurrentIndex = (lb.CurrentIndex + 1) % len(lb.Nodes)
		if lb.CurrentIndex == startIndex {
			return nil, fmt.Errorf("No available nodes to handle request")
		}
	}
}

func (lb *LoadBalancer) HandleRequest(requestSize int) error {
	node, err := lb.SelectNode(requestSize)
	if err != nil {
		return err
	}

	node.HandleRequest(requestSize)
	fmt.Printf("Request of size %d handled by node %s\n", requestSize, node.Address)
	return nil
}
