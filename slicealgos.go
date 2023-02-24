package gosort

import "golang.org/x/exp/constraints"

func CloneSlice[T any](v []T) []T {
	cpy := make([]T, len(v))
	copy(cpy, v)
	return cpy
}

func SlicesAreReversed[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	i := 0
	j := len(a) - 1
	for i <= j {
		if a[i] != b[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func Reverse[T constraints.Ordered](buffer []T) {
	i := 0
	j := len(buffer) - 1
	for i < j {
		buffer[i], buffer[j] = buffer[j], buffer[i]
		i++
		j--
	}
}

func IsSorted[T constraints.Ordered](v []T) bool {
	n := len(v)
	for i := n - 1; i > 0; i-- {
		if v[i] < v[i-1] {
			return false
		}
	}

	return true
}

func SlicesAreEqual[T comparable](a, b []T) bool {
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

type Predicate[T any] func(elem T) bool

func ForAll[T any](a []T, test Predicate[T]) bool {
	for _, v := range a {
		if !test(v) {
			return false
		}
	}

	return true
}

func IsGreaterThan[T constraints.Ordered](pivot T) Predicate[T] {
	return func(elem T) bool {
		return elem > pivot
	}
}

func IsLessThan[T constraints.Ordered](pivot T) Predicate[T] {
	return func(elem T) bool {
		return elem < pivot
	}
}
