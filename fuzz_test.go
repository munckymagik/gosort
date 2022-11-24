package gosort_test

import (
	"sort"
	"testing"

	"github.com/munckymagik/gosort"
)

type ByteSlice []byte

func (x ByteSlice) Len() int           { return len(x) }
func (x ByteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x ByteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func FuzzSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, stuff []byte) {
		given := make([]byte, len(stuff))
		copy(given, stuff)

		gosort.QuickSort(stuff)
		if !sort.IsSorted(ByteSlice(stuff)) {
			t.Errorf("Input %v, output %v", given, stuff)
		}
	})
}
