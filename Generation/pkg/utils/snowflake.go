package utils

import (
	"errors"
	"sync"
	"time"
)

// Constants for Snowflake algorithm
const (
	NodeBits  = 10
	StepBits  = 12
	NodeMax   = -1 ^ (-1 << NodeBits)
	StepMax   = -1 ^ (-1 << StepBits)
	TimeShift = NodeBits + StepBits
	NodeShift = StepBits
	Epoch     = 1704067200000 // 2024-01-01 00:00:00 UTC
)

// Node represents a Snowflake node
type Node struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

// NewNode creates a new Snowflake node
func NewNode(node int64) (*Node, error) {
	if node < 0 || node > NodeMax {
		return nil, errors.New("node number must be between 0 and 1023")
	}

	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

// Generate creates a unique ID
func (n *Node) Generate() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixMilli()

	if now < n.timestamp {
		now = n.timestamp // Clock moved backwards, refuse to generate id
	}

	if now == n.timestamp {
		n.step = (n.step + 1) & StepMax
		if n.step == 0 {
			for now <= n.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		n.step = 0
	}

	n.timestamp = now

	return ((now - Epoch) << TimeShift) | (n.node << NodeShift) | n.step
}
