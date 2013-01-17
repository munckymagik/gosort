package main

import (
    "testing"
    "sort" // The real 'sort' package from the standard library
)

/*
 * Check that the insertion sort implementation does not fail if an empty
 * input is passed.
 */
func TestBubbleSortShouldHandleEmptyInput(t *testing.T) {
    var a []int = make([]int, 0)
    BubbleSort(a)
}

/*
 * Check that the insertion sort implementation does not fail if a single
 * element input is passed.
 */
func TestBubbleSortShouldHandleOneElement(t *testing.T) {
    var a []int = make([]int, 1)
    BubbleSort(a)
}

/*
 * Check that the insertion sort implementation correctly processes an
 * input that is already sorted.
 */
func TestBubbleSortShouldHandleSorted(t *testing.T) {
    a := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9 }
    
    BubbleSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Corrupted already sorted array: ", a)
    }
}

/*
 * Check that the insertion sort implementation correctly processes an
 * input that is reverse sorted.
 */
func TestBubbleSortShouldHandleReverseSorted(t *testing.T) {
    a := []int{ 9, 8, 7, 6, 5, 4, 3, 2, 1 }
    
    BubbleSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Failed to sort reverse sorted array: ", a)
    }
}

/*
 * Check that the insertion sort implementation correctly processes
 * randomized inputs of various sizes.
 */
func TestBubbleSortShouldSortRandomized(t *testing.T) {
    a := []int{ 3, 2, 6, 5, 4, 8, 7, 1, 9 }
    
    BubbleSort(a)
    
    if !sort.IntsAreSorted(a) {
        t.Error("Failed to sort reverse sorted array: ", a)
    }
}
