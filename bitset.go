package bitset

import (
	"math/bits"
	"sync"
)

const mask = 0x3F

// Bitset -
type Bitset struct {
	mu   sync.RWMutex
	data []uint64
}

// NewBitset -
func NewBitset() *Bitset {
	return &Bitset{data: make([]uint64, 0)}
}

// Add -
func (b *Bitset) Add(value uint32) {
	index := getIndex(value)
	b.mu.Lock()
	for index >= len(b.data) {
		b.data = append(b.data, 0)
	}
	b.data[index] |= 1 << (value & mask)
	b.mu.Unlock()
}

// Remove -
func (b *Bitset) Remove(value uint32) {
	index := getIndex(value)
	b.mu.Lock()
	if index > len(b.data) {
		b.mu.Unlock()
		return
	}
	b.data[index] ^= 1 << (value & mask)
	b.mu.Unlock()
}

// Clear -
func (b *Bitset) Clear() {
	b.mu.Lock()
	b.data = b.data[:0]
	b.mu.Unlock()
}

// Values -
func (b *Bitset) Values() []uint32 {
	result := make([]uint32, 0)
	b.mu.RLock()
	for i, value := range b.data {
		if value == 0 {
			continue
		}

		for j := bits.TrailingZeros64(value); j <= bits.Len64(value); j++ {
			if value&(1<<uint(j)) > 0 {
				result = append(result, uint32(i*64+j))
			}
		}
	}
	b.mu.RUnlock()
	return result
}

func getIndex(value uint32) int {
	return int(value >> 6)
}
