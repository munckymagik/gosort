package gosort

import (
	"fmt"
	"testing"
	"testing/quick"

	"golang.org/x/exp/constraints"
)

func TestSortingAlgos(t *testing.T) {
	algos := []struct {
		sortInts    Sorter[int]
		sortStrings Sorter[string]
	}{
		{BubbleSort[int], BubbleSort[string]},
		{InsertionSort[int], InsertionSort[string]},
		{MergeSort[int], MergeSort[string]},
		{QuickSort[int], QuickSort[string]},
		{reverseMinHeapSort[int], reverseMinHeapSort[string]},
	}

	for _, algo := range algos {
		t.Run("Test basic properties", func(t *testing.T) {
			for desc, fixture := range fixtures() {
				t.Run(fmt.Sprint(desc), func(t *testing.T) {
					algo.sortInts(fixture)
					assertTrue(t, IsSorted(fixture))
				})
			}
		})
		t.Run("Property test int sorter", func(t *testing.T) {
			propertyTest(algo.sortInts, t)
		})
		t.Run("Property test string sorter", func(t *testing.T) {
			propertyTest(algo.sortStrings, t)
		})
	}
}

func fixtures() map[string][]int {
	return map[string][]int{
		"Should Handle Nil":            nil,
		"Should Handle Empty":          {},
		"Should Handle One Element":    {0},
		"Should Handle Already Sorted": {1, 2, 3, 4, 5},
		"Should Sort Reversed":         {5, 4, 3, 2, 1},
		"Should Sort Shuffled":         {4, 5, 3, 1, 2},
		"Should Sort Duplicates":       {6, 6, 4, 3, 8, 8, 3, 0, 3, 8},
	}
}

func propertyTest[T constraints.Ordered](sorter Sorter[T], t *testing.T) {
	alwaysSorts := func(input []T) bool {
		cpy := Clone(input)
		sorter(cpy)
		return IsSorted(cpy)
	}

	assertNoError(t, quick.Check(alwaysSorts, &quick.Config{MaxCount: 1000}))
}

func reverseMinHeapSort[T constraints.Ordered](a []T) {
	// MinHeap naturally reverse sorts its input. In order to pass these
	// tests we simply need to reverse the input after sorting.
	MinHeapSort(a)
	Reverse(a)
}
