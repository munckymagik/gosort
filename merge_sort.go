package gosort

import "golang.org/x/exp/constraints"

// MergeSort sorts an array of ordered elements, in place, using the Merge Sort algorithm.
func MergeSort[T constraints.Ordered](a []T) {
	mergeSort(a, 0, len(a)-1)
}

func mergeSort[T constraints.Ordered](a []T, start, end int) {
	if start < end {
		middle := (start + end) / 2

		mergeSort(a, start, middle)
		mergeSort(a, middle+1, end)

		merge(a, start, middle, end)
	}
}

func merge[T constraints.Ordered](a []T, start, middle, end int) {
	buffer := make([]T, end-start+1)
	copy(buffer, a[start:end+1])

	index := start
	left := 0
	right := middle + 1 - start
	leftEnd := right - 1
	rightEnd := end - start

	for left <= leftEnd && right <= rightEnd {
		if buffer[left] <= buffer[right] {
			a[index] = buffer[left]
			left++
		} else {
			a[index] = buffer[right]
			right++
		}
		index++
	}

	// If the right hand buffer is exhausted first then we need to copy over
	// the remaining left hand elements. If the left hand buffer is exhausted
	// instead then we don't need to do anything as the the right hand elements
	// will already be in the destination array.
	remainder := leftEnd + 1 - left

	if remainder > 0 {
		copy(a[index:], buffer[left:leftEnd+1])
	}
}
