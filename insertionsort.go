package gosort

import "golang.org/x/exp/constraints"

// InsertionSort sorts an array of ordered elements, in place, using the insertion sort algorithm.
func InsertionSort[T constraints.Ordered](a []T) {
	// Iterate up from the second element to the last
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		// Iterate down through the sorted subset [j-1 ... 0], shifting
		// elements upwards until key's position is found
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}

		// Insert key into its position in the sorted subset
		a[i+1] = key
	}
}
