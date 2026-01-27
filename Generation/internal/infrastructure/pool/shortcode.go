package pool

import (
	"context"
	"sync/atomic"
	"time"

	"go-link/common/pkg/datastructs/queue"
	"go-link/common/pkg/encoding"
	"go-link/common/pkg/unique"
	"go-link/generation/global"

	"go.uber.org/zap"
)

const (
	// DefaultCapacity is the default capacity of the short code pool.
	DefaultCapacity = 120_000

	// DefaultRefillInterval is the default interval to check for refill.
	DefaultRefillInterval = 1 * time.Second
)

// Options contains configuration for ShortCode pool.
type Options struct {
	Capacity       int
	RefillInterval time.Duration
}

// Option is a function that configures Options.
type Option func(*Options)

// WithCapacity sets the pool capacity.
func WithCapacity(capacity int) Option {
	return func(o *Options) {
		o.Capacity = capacity
	}
}

// WithRefillInterval sets the interval for refill checks.
func WithRefillInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.RefillInterval = interval
	}
}

// ShortCode holds a pre-generated pool of short codes for fast retrieval.
// Uses MPMC Queue for thread-safe, lock-free access.
type ShortCode struct {
	pool    *queue.MPMC[string]
	node    *unique.SnowflakeNode
	options *Options
	running atomic.Bool
	stopCh  chan struct{}
}

// NewShortCode creates a new pool with the given Snowflake node and options.
func NewShortCode(node *unique.SnowflakeNode, opts ...Option) *ShortCode {
	options := &Options{
		Capacity:       DefaultCapacity,
		RefillInterval: DefaultRefillInterval,
	}

	for _, opt := range opts {
		opt(options)
	}

	return &ShortCode{
		pool:    queue.NewMPMC[string](options.Capacity),
		node:    node,
		options: options,
		stopCh:  make(chan struct{}),
	}
}

// Start begins the background refill worker.
// It fills the pool initially and continuously refills.
func (p *ShortCode) Start(ctx context.Context) {
	if p.running.Swap(true) {
		return // Already running
	}

	// Initial fill
	// p.fillToCapacity()

	// Background refill worker
	go p.refillWorker(ctx)
}

// Stop gracefully stops the background worker.
func (p *ShortCode) Stop() {
	if !p.running.Swap(false) {
		return
	}
	close(p.stopCh)
}

// Get retrieves a pre-generated short code from the pool.
// Returns empty string and false if pool is exhausted.
func (p *ShortCode) Get() (string, bool) {
	return p.pool.Dequeue()
}

// GetOrGenerate retrieves from pool, or generates on-demand if pool is empty.
// This ensures the caller always gets a valid code.
func (p *ShortCode) GetOrGenerate() string {
	if code, ok := p.pool.Dequeue(); ok {
		global.LoggerZap.Info("Short code retrieved from pool", zap.String("shortCode", code), zap.Int64("poolSize", p.pool.Size()))
		return code
	}

	// Fallback: Generate on-demand (slower path)
	global.LoggerZap.Warn("Short code pool is empty, generating on-demand", zap.Int64("poolSize", p.pool.Size()))
	return encoding.Base62Encode(p.node.Generate())
}

// Size returns the current number of codes in the pool.
func (p *ShortCode) Size() int64 {
	return p.pool.Size()
}

// Capacity returns the maximum pool capacity.
func (p *ShortCode) Capacity() int {
	return p.options.Capacity
}

// IsRunning returns true if the background worker is running.
func (p *ShortCode) IsRunning() bool {
	return p.running.Load()
}

// refillWorker continuously monitors and refills the pool.
func (p *ShortCode) refillWorker(ctx context.Context) {
	ticker := time.NewTicker(p.options.RefillInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		case <-ticker.C:
			if !p.pool.IsFull() {
				p.fillToCapacity()
			}
		}
	}
}

// fillToCapacity attempts to fill the pool until it is full.
func (p *ShortCode) fillToCapacity() {
	for {
		// Optimization: Check size before generating
		if p.pool.IsFull() {
			return
		}

		id := p.node.Generate()
		code := encoding.Base62Encode(id)

		if !p.pool.Enqueue(code) {
			// Pool is full
			return
		}
	}
}
