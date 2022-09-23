package gosort

import "testing"

func TestHeapSortInvariant(t *testing.T) {
	for _, fixture := range fixtures {
		// clone the fixture to prevent accidental modification
		var cloned_fixture []int = make([]int, len(fixture))
		copy(cloned_fixture, fixture)

		heap := NewMinHeap(cloned_fixture)

		if !heap.checkInvariant() {
			t.Error("Fails invariant")
		}
	}
}
