package gosort

import (
	"testing"
)

func TestQuickSortPivotChoosers(t *testing.T) {
	pivotChoosers := []PivotChooser[int]{
		ChooseFirstElementPivot[int],
		ChooseLastElementPivot[int],
		ChooseMiddleElementPivot[int],
		NewChooseRandomElementPivot[int](1),
		ChooseMedianElementPivot[int],
	}

	for _, chooser := range pivotChoosers {
		curriedQuicksort := func(a []int) {
			QuickSortWithPivotChoice(a, chooser)
		}

		testSorting(curriedQuicksort, t)
	}
}

func TestPartition(t *testing.T) {
	a := []int{3, 8, 2, 5, 1, 4, 7, 6}

	partition(a, 0)

	success := a[2] == 3
	success = success && isTrueForAll(a[:2], createLessThanTest(3))
	success = success && isTrueForAll(a[3:], createGreaterThanTest(3))

	if !success {
		t.Errorf("Failed %v", a)
	}
}

func TestPartition2(t *testing.T) {
	a := []int{3, 8, 2, 5, 1, 4, 7, 6}

	partition(a, 3)

	success := a[4] == 5
	success = success && isTrueForAll(a[:4], createLessThanTest(5))
	success = success && isTrueForAll(a[5:], createGreaterThanTest(5))

	if !success {
		t.Errorf("Failed %v", a)
	}
}

type testFunc func(elem int) bool

func isTrueForAll(a []int, test testFunc) bool {

	for _, v := range a {
		if !test(v) {
			return false
		}
	}

	return true
}

func createGreaterThanTest(pivot int) testFunc {
	return func(elem int) bool {
		return elem > pivot
	}
}

func createLessThanTest(pivot int) testFunc {
	return func(elem int) bool {
		return elem < pivot
	}
}
