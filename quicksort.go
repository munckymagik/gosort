package gosort

import (
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

/*
 ChoosePivot defines the signature for a function that given a slice returns
 an index that can be used a pivot in the partitioning stage of QuickSort
 algorithm.
*/
type ChoosePivot[T constraints.Ordered] func(a []T) int

/*
 QuickSort sorts an array of integers, in place, using the QuickSort
 algorithm.
 This implementation uses random pivot selection.
*/
func QuickSort[T constraints.Ordered](a []T) {
	QuickSortWithPivotChoice(a, NewChooseRandomElementPivot[T](time.Now().UTC().UnixNano()))
}

/*
 QuickSortWithPivotChoice sorts an array of integers, in place, using the
 QuickSort algorithm, but also allows calling code to select how pivots
 are chosen. Pass one of the Choose<strategy>Pivot functions as the
 second argument to specify A QuickSort pivot selection strategy.
*/
func QuickSortWithPivotChoice[T constraints.Ordered](a []T, choosePivot ChoosePivot[T]) {
	if len(a) <= 1 {
		return
	}

	pivot := choosePivot(a)
	pivot = partition(a, pivot)
	QuickSortWithPivotChoice(a[:pivot], choosePivot)
	QuickSortWithPivotChoice(a[pivot+1:], choosePivot)
}

// A QuickSort pivot selection strategy: choose the first element. Use with
// QuickSortWithPivotChoice.
func ChooseFirstElementPivot[T constraints.Ordered](a []T) int {
	return 0
}

// A QuickSort pivot selection strategy: choose the last element. Use with
// QuickSortWithPivotChoice.
func ChooseLastElementPivot[T constraints.Ordered](a []T) int {
	return len(a) - 1
}

// A QuickSort pivot selection strategy: choose the middle element. Use with
// QuickSortWithPivotChoice.
func ChooseMiddleElementPivot[T constraints.Ordered](a []T) int {
	return (len(a) - 1) / 2
}

// A QuickSort pivot selection strategy: choose a random element. Use with
// QuickSortWithPivotChoice.
// Deprecated: this no longer sets the global seed so it will become deterministic.
// use NewChooseRandomElementPivot instead.
func ChooseRandomElementPivot[T constraints.Ordered](a []T) int {
	return rand.Intn(len(a))
}

// NewChooseRandomElementPivot returns a QuickSort pivot selection strategy: choose
// a random element. Use with QuickSortWithPivotChoice.
func NewChooseRandomElementPivot[T constraints.Ordered](seed int64) func([]T) int {
	r := rand.New(rand.NewSource(seed))
	return func(a []T) int {
		return r.Intn(len(a))
	}
}

/*
 A QuickSort pivot selection strategy: try to choose a median-ish element.
 Use with QuickSortWithPivotChoice.

 First selects the first, middle and last elements, then chooses
 the median of these three values and returns its index.
*/
func ChooseMedianElementPivot[T constraints.Integer | constraints.Float](a []T) int {
	first := 0
	middle := (len(a) - 1) / 2
	last := (len(a) - 1)

	median := last

	if ((a[first] - a[middle]) * (a[last] - a[first])) >= 0 {
		median = first
	} else if ((a[middle] - a[first]) * (a[last] - a[middle])) >= 0 {
		median = middle
	}

	return median
}

func partition[T constraints.Ordered](a []T, pivot int) int {
	// If the pivot isn't the first element, swap it so it is
	if pivot > 0 {
		a[0], a[pivot] = a[pivot], a[0]
	}

	p := a[0]
	i := 1
	for j := i; j < len(a); j++ {
		if a[j] < p {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	// Put the pivot element into position
	a[0], a[i-1] = a[i-1], a[0]

	return i - 1
}
