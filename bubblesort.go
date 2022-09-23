package gosort

import "golang.org/x/exp/constraints"

// BubbleSort sorts an array of ordered elements, in place, using the bubble sort algorithm.
func BubbleSort[T constraints.Ordered](a []T) {
	// Iterate from 0 up to the second to last element
	for i := 0; i < len(a)-1; i++ {
		// Iterate down from the last element to element i+1
		for j := len(a) - 1; j > i; j-- {
			// Compare each pair of elements on the way
			if a[j] < a[j-1] {
				// Swap the elements if they are out of sort order
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
