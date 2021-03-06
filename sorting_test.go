package gosort

import (
    "sort" // The real 'sort' package from the standard library
    "testing"
)

var fixtures = map[string][]int{
    "Should Handle Nil":            nil,
    "Should Handle Empty":          {},
    "Should Handle One Element":    {0},
    "Should Handle Already Sorted": {1, 2, 3, 4, 5},
    "Should Sort Reversed":         {5, 4, 3, 2, 1},
    "Should Sort Shuffled":         {4, 5, 3, 1, 2},
    "Should Sort Duplicates":       {6, 6, 4, 3, 8, 8, 3, 0, 3, 8},
}

func runTests(alg func([]int), t *testing.T) {
    for desc, fixture := range fixtures {
        // clone the fixture to prevent accidental modification
        var cloned_fixture []int = make([]int, len(fixture))
        copy(cloned_fixture, fixture)

        // Execute in a separate function in order to handle errors
        func() {
            // Handle errors
            defer func() {
                if err := recover(); err != nil {
                    t.Errorf("Error occured in '%s'. Err: %s, in: %+v, out: %+v",
                        err, desc, fixture, cloned_fixture)
                }
            }()

            // Now execute the test
            alg(cloned_fixture)

            // Check that the operation was successful
            if !sort.IntsAreSorted(cloned_fixture) {
                t.Errorf("Test '%s' failed. In: %v, out: %v",
                    desc, fixture, cloned_fixture)
            }
        }()
    }
}

func reverse(buffer []int) {
    start := 0
    end := len(buffer) - 1
    for start < end {
        buffer[start], buffer[end] = buffer[end], buffer[start]
        start++
        end--
    }
}

func TestBubbleSort(t *testing.T) {
    runTests(BubbleSort, t)
}

func TestInsertionSort(t *testing.T) {
    runTests(InsertionSort, t)
}

func TestMergeSort(t *testing.T) {
    runTests(MergeSort, t)
}

func TestQuickSort(t *testing.T) {
    runTests(QuickSort, t)
}

func TestQuickSortPivotChoosers(t *testing.T) {
    pivotChoosers := []ChoosePivot{
        ChooseFirstElementPivot,
        ChooseLastElementPivot,
        ChooseMiddleElementPivot,
        ChooseRandomElementPivot,
    }

    for _, chooser := range pivotChoosers {
        curriedQuicksort := func(a []int) {
            QuickSortWithPivotChoice(a, chooser)
        }

        runTests(curriedQuicksort, t)
    }
}

func TestMinHeapSort(t *testing.T) {
    // MinHeap naturally reverse sorts its input. In order to pass these
    // tests we simply need to reverse the input after sorting.
    forwardSort := func(a []int) {
        MinHeapSort(a)
        reverse(a)
    }

    runTests(forwardSort, t)
}

func TestHeapSortInvariant(t *testing.T) {
    for _, fixture := range fixtures {
        // clone the fixture to prevent accidental modification
        var cloned_fixture []int = make([]int, len(fixture))
        copy(cloned_fixture, fixture)

        heap := NewMinHeap(cloned_fixture)

        if !heap.checkInvariant() {
            t.Error("Fails invariant")
        }
    }
}
