package gosort_test

import (
	"testing"

	"github.com/munckymagik/gosort"
)

func assertSlicesEq(t *testing.T, got, expected []int) {
	if len(got) != len(expected) {
		t.Errorf("Test failed. Not same length. Got: %v, expected: %v", got, expected)
	}

	for i := 0; i < len(expected); i++ {
		if got[i] != expected[i] {
			t.Errorf("Test failed. Not equal. Got: %v, expected: %v", got, expected)
		}
	}
}

func TestBubbleSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.BubbleSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

func TestInsertionSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.InsertionSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

func TestMergeSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.MergeSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

func TestQuickSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{1, 2, 3}

	// When
	gosort.QuickSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}

func TestReverseMinHeapSortInts(t *testing.T) {
	// Given
	input := []int{3, 1, 2}
	expected := []int{3, 2, 1}

	// When
	gosort.MinHeapSort(input)

	// Then
	assertSlicesEq(t, input, expected)
}
