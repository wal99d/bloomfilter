// Package bloomfilter implements a simple bloom filter using hashing functions.
package bloomfilter

import (
	"hash/fnv"
	"sync"
)

// BloomFilter is a struct representing a bloom filter.
type BloomFilter struct {
	size   int
	bits   []bool
	funcs  []HashFunc
	rwlock sync.RWMutex
}

// HashFunc is a type representing a hashing function for the bloom filter.
type HashFunc func(string) uint32

// NewBloomFilter returns a new bloom filter instance.
//
// The size parameter specifies the size of the in-memory bloom filter array.
// The funcs parameter is a slice of HashFuncs that will be used to hash inputs to the bloom filter.
func NewBloomFilter(size int, funcs []HashFunc) *BloomFilter {
	b := &BloomFilter{
		size:   size,
		bits:   make([]bool, size),
		funcs:  funcs,
		rwlock: sync.RWMutex{},
	}
	return b
}

// Add adds a new input string to the bloom filter.
func (b *BloomFilter) Add(key string) {
	b.rwlock.Lock()
	defer b.rwlock.Unlock()

	for _, f := range b.funcs {
		index := f(key) % uint32(b.size)
		b.bits[index] = true
	}
}

// Remove removes an input string from the bloom filter
func (b *BloomFilter) Remove(key string) {
	b.rwlock.Lock()
	defer b.rwlock.Unlock()

	for _, f := range b.funcs {
		index := f(key) % uint32(b.size)
		b.bits[index] = false
	}
}

// Check checks whether an input string is probably in the bloom filter.
func (b *BloomFilter) Check(key string) bool {
	b.rwlock.RLock()
	defer b.rwlock.RUnlock()

	for _, f := range b.funcs {
		index := f(key) % uint32(b.size)
		if !b.bits[index] {
			return false
		}
	}
	return true
}

// NewHashFunc returns a new FNV-1 hashing function for the bloom filter.
func NewHashFunc() HashFunc {
	return func(key string) uint32 {
		h := fnv.New32a()
		h.Write([]byte(key))
		return h.Sum32()
	}
}
