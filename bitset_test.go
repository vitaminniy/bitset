package bitset

import (
	"fmt"
	"testing"
)

func Test_Values(t *testing.T) {
	valueSets := [][]uint32{
		[]uint32{1, 2, 3},
		[]uint32{},
		[]uint32{0},
		[]uint32{100, 4, 7, 13, 89, 88, 92, 1},
	}

	for _, values := range valueSets {
		t.Run(fmt.Sprintf("Values %d", len(values)), func(t *testing.T) {
			bs := NewBitset()
			for _, value := range values {
				bs.Add(value)
			}

			actual := bs.Values()
			if len(actual) != len(values) {
				t.Fatalf("Actual length (%d) must be equal to expected length (%d)", len(actual), len(values))
			}

			for i := range values {
				var found bool
				for j := range actual {
					if values[i] == actual[j] {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("Item %d doesn't exist in actual result", values[i])
				}
			}
		})
	}
}

func Benchmark_Add(b *testing.B) {
	values := []uint32{10e0, 10e1, 10e2, 10e3, 10e4, 10e5}
	for _, value := range values {
		b.Run(fmt.Sprintf("Add %d items", value), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bs := NewBitset()
				for j := uint32(0); j < value; j++ {
					bs.Add(j)
				}
			}
		})
	}
}
