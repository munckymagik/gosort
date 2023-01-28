package gosort

import (
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
		testSorting(algo.sortInts, t)
		propertyTest(algo.sortInts, t)
		propertyTest(algo.sortStrings, t)
	}
}

func propertyTest[T constraints.Ordered](sorter Sorter[T], t *testing.T) {
	alwaysSorts := func(input []T) bool {
		cpy := clone(input)
		sorter(cpy)
		return isSorted(cpy)
	}

	cfg := quick.Config{
		MaxCount: 1000,
	}
	if err := quick.Check(alwaysSorts, &cfg); err != nil {
		t.Error(err)
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

func testSorting(alg func([]int), t *testing.T) {
	for desc, fixture := range fixtures() {
		// Execute in a separate function in order to handle errors
		func() {
			// Handle errors
			defer func() {
				if err := recover(); err != nil {
					t.Errorf("Error occured in '%s'. Err: %s, in: %+v, out: %+v",
						err, desc, fixture, fixture)
				}
			}()

			// Now execute the test
			alg(fixture)

			// Check that the operation was successful
			if !isSorted(fixture) {
				t.Errorf("Test '%s' failed. In: %v, out: %v",
					desc, fixture, fixture)
			}
		}()
	}
}

func reverseMinHeapSort[T constraints.Ordered](a []T) {
	// MinHeap naturally reverse sorts its input. In order to pass these
	// tests we simply need to reverse the input after sorting.
	MinHeapSort(a)
	reverse(a)
}

func reverse[T constraints.Ordered](buffer []T) {
	start := 0
	end := len(buffer) - 1
	for start < end {
		buffer[start], buffer[end] = buffer[end], buffer[start]
		start++
		end--
	}
}

func isSorted[T constraints.Ordered](v []T) bool {
	n := len(v)
	for i := n - 1; i > 0; i-- {
		if v[i] < v[i-1] {
			return false
		}
	}

	return true
}

func clone[T any](v []T) []T {
	cpy := make([]T, len(v))
	copy(cpy, v)
	return cpy
}
