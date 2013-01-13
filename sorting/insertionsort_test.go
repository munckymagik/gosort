package insertionsort

import (
    "testing"
    "sort" // The real 'sort' package from the standard library
)

/*
 * Check that the insertion sort implementation does not fail if an empty
 * input is passed.
 */
func TestInsertionSortShouldHandleEmptyInput(t *testing.T) {
    var a []int = make([]int, 0)
    InsertionSort(a)
}

/*
 * Check that the insertion sort implementation does not fail if a single
 * element input is passed.
 */
func TestInsertionSortShouldHandleOneElement(t *testing.T) {
    var a []int = make([]int, 1)
    InsertionSort(a)
}

/*
 * Check that the insertion sort implementation correctly processes an
 * input that is already sorted.
 */
func TestInsertionSortShouldHandleSorted(t *testing.T) {
    a := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }
    
    InsertionSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Corrupted already sorted array: ", a)
    }
}

/*
 * Check that the insertion sort implementation correctly processes an
 * input that is reverse sorted.
 */
func TestInsertionSortShouldHandleReverseSorted(t *testing.T) {
    a := []int{ 9, 8, 7, 6, 5, 4, 3, 2, 1 }
    
    InsertionSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Failed to sort reverse sorted array: ", a)
    }
}

/*
 * Check that the insertion sort implementation correctly processes
 * randomized inputs of various sizes.
 */
func TestInsertionSortShouldSortRandomized(t *testing.T) {
    a := []int{ 3, 2, 6, 5, 4, 8, 7, 1, 9 }
    
    InsertionSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Failed to sort reverse sorted array: ", a)
    }
}
