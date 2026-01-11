package unique

import (
	"errors"
	"sync"

	"go-link/common/pkg/settings"
	t "go-link/common/pkg/timer"
)

// Node represents a Snowflake node
type SnowflakeNode struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64

	// Configuration
	epoch    int64
	nodeBits uint8
	stepBits uint8

	// Pre-calculated masks and shifts
	nodeMax   int64
	stepMax   int64
	timeShift uint8
	nodeShift uint8

	// Dependencies
	clock t.Timer
}

func NewSnowflakeNode(config settings.SnowflakeNode, clock t.Timer) (*SnowflakeNode, error) {
	nodeMax := int64(-1 ^ (-1 << config.Config.Node))
	stepMax := int64(-1 ^ (-1 << config.Config.Step))

	if config.WorkerID < 0 || config.WorkerID > nodeMax {
		return nil, errors.New("node ID exceeds maximum allowed by configuration")
	}

	return &SnowflakeNode{
		timestamp: 0,
		node:      config.WorkerID,
		step:      0,

		epoch:    config.Config.Epoch,
		nodeBits: config.Config.Node,
		stepBits: config.Config.Step,

		nodeMax:   nodeMax,
		stepMax:   stepMax,
		timeShift: config.Config.Node + config.Config.Step,
		nodeShift: config.Config.Step,

		clock: clock,
	}, nil
}

// Generate creates a unique ID
func (n *SnowflakeNode) Generate() int64 {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := n.clock.Now().UnixMilli()

	if now < n.timestamp {
		now = n.timestamp
	}

	if now == n.timestamp {
		n.step = (n.step + 1) & n.stepMax
		if n.step == 0 {
			for now <= n.timestamp {
				now = n.clock.Now().UnixMilli()
			}
		}
	} else {
		n.step = 0
	}

	n.timestamp = now

	return ((now - n.epoch) << n.timeShift) | (n.node << n.nodeShift) | n.step
}
