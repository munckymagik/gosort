package gosort

import (
	"testing"
	"testing/quick"
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
		itSorts := func(v []int) bool {
			cpy := CloneSlice(v)
			QuickSortWithPivotChoice(cpy, chooser)
			return IsSorted(cpy)
		}
		assertNoError(t, quick.Check(itSorts, nil))
	}
}

func TestPartition(t *testing.T) {
	a := []int{3, 8, 2, 5, 1, 4, 7, 6}

	partition(a, 0)

	assertTrue(t, a[2] == 3)
	assertTrue(t, ForAll(a[:2], IsLessThan(3)))
	assertTrue(t, ForAll(a[3:], IsGreaterThan(3)))
}

func TestPartition2(t *testing.T) {
	a := []int{3, 8, 2, 5, 1, 4, 7, 6}

	partition(a, 3)

	assertTrue(t, a[4] == 5)
	assertTrue(t, ForAll(a[:4], IsLessThan(5)))
	assertTrue(t, ForAll(a[5:], IsGreaterThan(5)))
}
