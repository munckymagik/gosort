package gosort

import (
	"sort" // The real 'sort' package from the standard library
	"testing"
)

func TestSortingAlgos(t *testing.T) {
	algos := []Sorter{
		BubbleSort,
		InsertionSort,
		MergeSort,
		QuickSort,
		reverseMinHeapSort,
	}

	for _, algo := range algos {
		testSorting(algo, t)
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
			if !sort.IntsAreSorted(fixture) {
				t.Errorf("Test '%s' failed. In: %v, out: %v",
					desc, fixture, fixture)
			}
		}()
	}
}

// func testGeneric(sort Sorter[string], t *testing.T) {
// 	a := []string{"bb", "0", "ba", "7", "aa"}
// 	expected := []string{"0", "7", "aa", "ba", "bb"}
// 	sort(a)
// 	if !reflect.DeepEqual(a, expected) {
// 		t.Errorf("Expected %v to equal %v", a, expected)
// 	}
// }

func reverseMinHeapSort(a []int) {
	// MinHeap naturally reverse sorts its input. In order to pass these
	// tests we simply need to reverse the input after sorting.
	MinHeapSort(a)
	reverse(a)
}

func reverse(buffer []int) {
	start := 0
	end := len(buffer) - 1
	for start < end {
		buffer[start], buffer[end] = buffer[end], buffer[start]
		start++
		end--
	}
}
