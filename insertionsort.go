package sorting

/*
 * Sort an array of integers, in place, using the insertion sort algorithm.
 */
func InsertionSort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}

		a[i+1] = key
	}
}

