package main

/*
 * Sort an array of integers, in place, using the bubble sort algorithm.
 */
func BubbleSort(a []int) {
    for i := 0; i < len(a) - 1; i++ {
        for j := len(a) - 1; j > i; j-- {
            if a[j] < a[j - 1] {
                a[j], a[j - 1] = a[j - 1], a[j]
            }
        }
    }
}