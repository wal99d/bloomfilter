package internals

import (
	"hash/fnv"
	"sync"
)

type BloomFilter struct {
	size   int
	bits   []bool
	funcs  []HashFunc
	rwlock sync.RWMutex
}

type HashFunc func(string) uint32

func NewBloomFilter(size int, funcs []HashFunc) *BloomFilter {
	b := &BloomFilter{
		size:   size,
		bits:   make([]bool, size),
		funcs:  funcs,
		rwlock: sync.RWMutex{},
	}
	return b
}

func (b *BloomFilter) Add(key string) {
	b.rwlock.Lock()
	defer b.rwlock.Unlock()

	for _, f := range b.funcs {
		index := f(key) % uint32(b.size)
		b.bits[index] = true
	}
}

func (b *BloomFilter) Remove(key string) {
	b.rwlock.Lock()
	defer b.rwlock.Unlock()

	for _, f := range b.funcs {
		index := f(key) % uint32(b.size)
		b.bits[index] = false
	}
}

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

func NewHashFunc() HashFunc {
	return func(key string) uint32 {
		h := fnv.New32a()
		h.Write([]byte(key))
		return h.Sum32()
	}
}
