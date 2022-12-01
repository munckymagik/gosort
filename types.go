package gosort

import "golang.org/x/exp/constraints"

// Sorter defines the signature for a function that sorts a slice of
// ordered elements in-place.
type Sorter[T constraints.Ordered] func(a []T)
