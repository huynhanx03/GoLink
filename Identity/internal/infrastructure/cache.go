package infrastructure

import (
	"go-link/common/pkg/common/cache/tinylfu"

	"go-link/identity/global"
)

const (
	cacheMaxCost     = 1000
	cacheNumCounters = 10000
	cacheBufferSize  = 64
)

// SetupCache initializes the cache.
func SetupCache() {
	config := tinylfu.Config{
		MaxCost:     cacheMaxCost,
		NumCounters: cacheNumCounters,
		BufferSize:  cacheBufferSize,
	}

	global.Tinylfu = tinylfu.New[string, any](config)
}
