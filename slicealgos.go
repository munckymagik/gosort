package gosort

import "golang.org/x/exp/constraints"

func clone[T any](v []T) []T {
	cpy := make([]T, len(v))
	copy(cpy, v)
	return cpy
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

func slicesAreEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
