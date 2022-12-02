package gosort_test

import (
	"testing"

	"github.com/munckymagik/gosort"
)

// ----------------------------------------------------------------------------
// Step 1: examine an already generic function
//
// BubbleSort has already been made generic.
//
// There are two tests for it here: one with an integer inputs and another with
// string inputs.
//
// Have a look at its implementation, then apply what you learned in Step 2.

func TestBubbleSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.BubbleSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

func TestBubbleSortStrings(t *testing.T) {
	// Given
	input := []string{"3", "1", "2"}
	expected := []string{"1", "2", "3"}

	// When
	gosort.BubbleSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

// ----------------------------------------------------------------------------
// Step 2: genericise a function
//
// InsertionSort is not yet generic and only supports integer inputs.
//
// Uncomment the TestInsertionSortStrings test case which uses string inputs,
// then copy what was done to BubbleSort to make InsertionSort generic and the
// test compile and pass.

func TestInsertionSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.InsertionSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

// UNCOMMENT AND MAKE THIS COMPILE AND PASS
// func TestInsertionSortStrings(t *testing.T) {
// 	// Given
// 	input := []string{"3", "1", "2"}
// 	expected := []string{"1", "2", "3"}

// 	// When
// 	gosort.InsertionSort(input)

// 	// Then
// 	assertSlicesEq(t, input, expected)
// }

// ----------------------------------------------------------------------------
// Step 3: genericise a group of related functions
//
// MergeSort is implemented using more than one function. You will need to
// adapt each component function to make the entire implementation generic.
//
// Uncomment the TestMergeSortStrings test case which uses string inputs, then
// use what you learned in the two previous exercises

func TestMergeSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.MergeSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

// UNCOMMENT AND MAKE THIS COMPILE AND PASS
// func TestMergeSortStrings(t *testing.T) {
// 	// Given
// 	input := []string{"3", "1", "2"}
// 	expected := []string{"1", "2", "3"}

// 	// When
// 	gosort.MergeSort(input)

// 	// Then
// 	assertSlicesEq(t, input, expected)
// }

// ----------------------------------------------------------------------------
// Step 4: genericise a type declaration
//
// QuickSort is implemented using more than one function just like MergeSort.
// But because it supports several different strategies for picking a pivot
// point, this time we will need to genercise the PivotChooser type declaration
// and the function which implement it.
//
// Uncomment the TestQuickSortStrings test case which uses string inputs, then
// work your way through until it compiles and passes.

func TestQuickSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.QuickSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

// UNCOMMENT AND MAKE THIS COMPILE AND PASS
// func TestQuickSortStrings(t *testing.T) {
// 	// Given
// 	input := []string{"3", "1", "2"}
// 	expected := []string{"1", "2", "3"}

// 	// When
// 	gosort.QuickSort(input)

// 	// Then
// 	assertSlicesEq(t, input, expected)
// }

// ----------------------------------------------------------------------------
// Step 5: genericise a structure
//
// MinHeapSort is implemented using a struct called MinHeap which holds its
// working state. So far we've been applying generics to functions. Now we will
// learn how to apply to structs and methods.
//
// Uncomment the TestReverseMinHeapSortStrings test case which uses string
// inputs, then work your way through until it compiles and passes.
func TestReverseMinHeapSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{3, 2, 1}

	// When
	gosort.MinHeapSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

// UNCOMMENT AND MAKE THIS COMPILE AND PASS
// func TestReverseMinHeapSortStrings(t *testing.T) {
// 	// Given
// 	input := []string{"3", "1", "2"}
// 	expected := []string{"3", "2", "1"}

// 	// When
// 	gosort.MinHeapSort(input)

// 	// Then
// 	assertSlicesEq(t, input, expected)
// }

// ----------------------------------------------------------------------------
// Step 6: further work
//
// These exercises were an introduction. But there are several further topics
// to understand. In particular:
//
// * What is type inference?
// * What are type sets?
// * What are union type constaints?
// * How do we define our own type constraints?
// * What is type constraint inference and why is it useful?
//
// Follow these two tutorials to find out:
// 1. https://go.dev/doc/tutorial/generics
// 2. https://go.dev/blog/intro-generics

// ----------------------------------------------------------------------------
// This is a helper function, ... oh look it's generic! In this case, we only
// do comparisons, so notice that we are able to specify the builtin
// "comparable" type constraint.

func assertSlicesEq[E comparable](t *testing.T, got, expected []E) {
	t.Helper()
	fail := func() {
		t.Helper()
		t.Errorf("Test failed. Not same length. Got: %v, expected: %v", got, expected)
	}

	if len(got) != len(expected) {
		fail()
	}

	for i := 0; i < len(expected); i++ {
		if got[i] != expected[i] {
			fail()
			return
		}
	}
}
