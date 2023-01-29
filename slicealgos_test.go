package gosort

import (
	"reflect"
	"sort"
	"testing"
	"testing/quick"
)

func TestSlicesAreEqual(t *testing.T) {
	t.Run("basic properties", func(t *testing.T) {
		assertTrue(t, SlicesAreEqual[int](nil, nil))
		assertTrue(t, SlicesAreEqual[string](nil, nil))
		assertTrue(t, !SlicesAreEqual(nil, []int{1}))
		assertTrue(t, !SlicesAreEqual([]int{1}, nil))
		assertTrue(t, SlicesAreEqual([]int{1}, []int{1}))
		assertTrue(t, !SlicesAreEqual([]int{1}, []int{2}))
	})

	t.Run("is consistent with reflect.DeepEqual", func(t *testing.T) {
		deepEq := func(a []int, b []int) bool { return reflect.DeepEqual(a, b) }
		assertNoError(t, quick.CheckEqual(SlicesAreEqual[int], deepEq, &quick.Config{MaxCount: 1000}))
	})
}

func TestClone(t *testing.T) {
	t.Run("when the input is nil it returns a nil slice", func(t *testing.T) {
		cpy := CloneSlice[int](nil)
		assertIsNilSlice(t, cpy)
	})
	t.Run("it clones the underlying array", func(t *testing.T) {
		a := []int{1, 2, 3}

		result := CloneSlice(a)

		assertSlicesEqual(t, a, result)
		result[0] = 0
		assertSlicesNotEqual(t, a, result)
	})
}

func TestReverse(t *testing.T) {
	t.Run("when the input is nil it does not panic", func(t *testing.T) {
		Reverse[float32](nil)
	})

	t.Run("is symmetric", func(t *testing.T) {
		isSymmetric := func(v []int) bool {
			copy := CloneSlice(v)

			Reverse(v)

			if len(v) > 1 && SlicesAreEqual(v, copy) {
				return false
			}

			Reverse(v)

			return SlicesAreEqual(v, copy)
		}
		assertNoError(t, quick.Check(isSymmetric, &quick.Config{MaxCount: 1000}))
	})
}

func TestIsSorted(t *testing.T) {
	t.Run("when the input is a nil slice it return true", func(t *testing.T) {
		assertTrue(t, IsSorted[float32](nil))
	})

	t.Run("is consistent with sort.SliceIsSorted", func(t *testing.T) {
		assertNoError(t, quick.CheckEqual(IsSorted[int], sort.IntsAreSorted, &quick.Config{MaxCount: 1000}))
	})
}
