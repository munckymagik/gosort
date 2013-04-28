package gosort

import (
    "testing"
)

type testFunc func(elem int) bool

func trueForAll(a []int, test testFunc) bool {
    
    for _, v := range a {
        if ! test(v) {
            return false
        }
    }
    
    return true
}

func genGreaterThan(pivot int) testFunc {
    return func(elem int) bool {
        return elem > pivot
    }
}

func genLessThan(pivot int) testFunc {
    return func(elem int) bool {
        return elem < pivot
    }
}

func TestPartition(t *testing.T) {
    a := []int{ 3, 8, 2, 5, 1, 4, 7, 6 }
    
    partition(a, 0)
    
    success := a[2] == 3
    success = success && trueForAll(a[:2], genLessThan(3))
    success = success && trueForAll(a[3:], genGreaterThan(3))
    
    if !success {
        t.Errorf("Failed %v", a)
    }
}

func TestPartition2(t *testing.T) {
    a := []int{ 3, 8, 2, 5, 1, 4, 7, 6 }
    
    partition(a, 3)
    
    success := a[4] == 5
    success = success && trueForAll(a[:4], genLessThan(5))
    success = success && trueForAll(a[5:], genGreaterThan(5))
    
    if !success {
        t.Errorf("Failed %v", a)
    }
}
