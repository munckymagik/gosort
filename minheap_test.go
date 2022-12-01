package gosort

import "testing"

func TestHeapSortInvariant(t *testing.T) {
	for _, fixture := range fixtures() {
		heap := NewMinHeap(fixture)

		if !heap.checkInvariant() {
			t.Error("Fails invariant")
		}
	}
}
