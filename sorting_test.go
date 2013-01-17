package main

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

func TestBubbleSort(t *testing.T) {
    runTests(BubbleSort, t)
}

func TestInsertionSort(t *testing.T) {
    runTests(InsertionSort, t)
}
