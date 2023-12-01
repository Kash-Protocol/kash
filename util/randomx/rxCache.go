package randomx

import "C"
import (
	"bytes"
)

// NewRxCache creates a new RxCache with the specified flags.
func NewRxCache(flags ...Flag) (*RxCache, error) {
	cache, err := AllocCache(flags...)
	if cache == nil {
		return nil, err
	}

	return &RxCache{cache: cache}, nil
}

// Close releases the resources associated with the RxCache.
func (c *RxCache) Close() {
	if c.cache != nil {
		ReleaseCache(c.cache)
	}
}

// Init initializes the RxCache with the given seed. Returns true if initialization was successful.
func (c *RxCache) Init(seed []byte) bool {
	if c.IsReady(seed) {
		return false
	}

	c.seed = seed
	InitCache(c.cache, c.seed)

	c.initCount++

	return true
}

// IsReady checks if the RxCache is ready and matches the given seed.
func (c *RxCache) IsReady(seed []byte) bool {
	return (c.initCount > 0) && (bytes.Compare(c.seed, seed) == 0)
}
