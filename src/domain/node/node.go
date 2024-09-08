package node

import "time"

type Node struct {
	Address         string
	BPMLimit        int
	RPMLimit        int
	CurrentBytes    int
	CurrentRequests int
	LastReset       time.Time
}

func NewNode(address string, bpmLimit, rpmLimit int) *Node {
	return &Node{
		Address:         address,
		BPMLimit:        bpmLimit,
		RPMLimit:        rpmLimit,
		CurrentBytes:    0,
		CurrentRequests: 0,
		LastReset:       time.Now(),
	}
}

func (n *Node) CanHandleRequest(requestSize int) bool {
	now := time.Now()

	if now.Sub(n.LastReset) >= time.Minute {
		n.CurrentBytes = 0
		n.CurrentRequests = 0
		n.LastReset = now
	}

	if n.CurrentBytes+requestSize > n.BPMLimit || n.CurrentRequests+1 > n.RPMLimit {
		return false
	}

	return true
}

func (n *Node) HandleRequest(requestSize int) {
	n.CurrentBytes += requestSize
	n.CurrentRequests++
}
