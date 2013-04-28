package gosort

func QuickSort(a []int) {
    if len(a) <= 1 {
        return
    }
    
    pivot := choosePivot(a)
    pivot = partition(a, pivot)
    QuickSort(a[:pivot])
    QuickSort(a[pivot+1:])
}

func choosePivot(a []int) int {
    // Simplest strategy: choose the first element
    return 0
}

func partition(a []int, pivot int) int {
    // If the pivot isn't the first element, swap it so it is
    if pivot > 0 {
        a[0], a[pivot] = a[pivot], a[0]
    }
    
    p := a[0]
    i := 1
    for j := i; j < len(a); j++ {
        if a[j] < p {
            a[i], a[j] = a[j], a[i]
            i++
        }
    }
    
    // Put the pivot element into position
    a[0], a[i-1] = a[i-1], a[0]
    
    return i-1
}

