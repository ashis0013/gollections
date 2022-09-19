package gollections

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Filters the slice based on the given predicate
func Filter[T any] (slice []T, predicate func(T) bool) []T {
    var filteredSlice = []T{}
    if predicate == nil { return filteredSlice }
    for _, elem := range slice {
        if predicate(elem) {
            filteredSlice = append(filteredSlice, elem)
        }
    }
    return filteredSlice
}

// Filters the slice based on the given predicate
// The predicate also considers index of each element
func FilterIndexed[T any] (slice []T, predicate func(int, T) bool) []T {
    var filteredSlice = []T{}
    if predicate == nil { return filteredSlice }
    for i, elem := range slice {
        if predicate(i, elem) {
            filteredSlice = append(filteredSlice, elem)
        }
    }
    return filteredSlice
}

// Applies the transform function on each element of the slice 
// and returns a slice of the transformed value
func Map[T any, R any] (slice []T, transform func(T) R) []R {
    var transformedSlice = []R{}
    if transform == nil { return transformedSlice }
    for _, elem := range slice {
        transformedSlice = append(transformedSlice, transform(elem))
    }
    return transformedSlice 
}

// Applies the transform function on each element of the slice 
// and returns a slice of the transformed value
// The transform also considers index of each element
func MapIndexed[T any, R any] (slice []T, transform func(int, T) R) []R {
    var transformedSlice = []R{}
    if transform == nil { return transformedSlice }
    for i, elem := range slice {
        transformedSlice = append(transformedSlice, transform(i, elem))
    }
    return transformedSlice 
}

//Returns true if all elements satisfy the given predicate
func All[T any] (slice []T, predicate func(T) bool) bool {
    if predicate == nil { return false }
    for _, elem := range slice {
        if !predicate(elem) {
            return false
        }
    }
    return true
} 

//Returns true if any one of elements satisfy the given predicate
func Any[T any] (slice []T, predicate func(T) bool) bool {
    if predicate == nil { return false }
    for _, elem := range slice {
        if predicate(elem) {
            return true
        }
    }
    return false
}

// Returns a map generated from the given slice, using the given transform
func Associate[T any, K comparable, V any] (slice []T, transform func(T) (K, V)) map[K]V {
    hashMap := make(map[K]V)
    if transform == nil { return hashMap }
    for _, elem := range slice {
        key, value := transform(elem)
        hashMap[key] = value
    }
    return hashMap
}

// Returns whether the given slice contains the given target value
func Contains[T comparable] (slice []T, target T) bool {
    for _, elem := range slice {
        if elem == target {
            return true
        }
    }
    return false
}

// Drops the value from the given slice at index and returns.
// The original slice remains unchanged.
func Drop[T any] (slice []T, index int) []T {
    if index < 0 || index > len(slice) {
        return slice
    }
    return FilterIndexed(slice, func(i int, elem T) bool {
        return i != index
    })
}

// Returns the first value the satisfies the given predicate.
// Raises error if there is no such element
func First[T any] (slice []T, predicate func(T) bool) (T, error) {
    if predicate == nil { return zero[T](), errors.New("nil function pointer passed")}
    for _, elem := range slice {
        if predicate(elem) {
            return elem, nil
        }
    }
    return zero[T](), errors.New("No element found satisfying predicate")
}

// Returns the first value the satisfies the given predicate.
// If no such element is there then returns defaultValue
func FirstOrDefault[T any] (slice []T, defaultValue T, predicate func(T) bool) T {
    first, err := First(slice, predicate)
    if err != nil {
        return defaultValue
    }
    return first
}

// Flattens a 2d slice into a 1d slice
func Flatten[T any] (slices [][]T) []T {
    var flattenedSlice []T
    for _, slice := range slices {
        for _, elem := range slice {
            flattenedSlice = append(flattenedSlice, elem)
        }
    }
    return flattenedSlice
}

// Accumulates value starting with the given initial 
// by performing operation on the slice from left to right
func Fold[T any, R any] (slice []T, initial R, operation func(T, R) R) R {
    accumulator := initial
    if operation == nil { return accumulator }
    for _, elem := range slice {
        accumulator = operation(elem, accumulator)
    }
    return accumulator
}

// Accumulates value starting with the given initial 
// by performing operation on the slice from left to right
// The operation also considers index of each element
func FoldIndexed[T any, R any] (slice []T, initial R, operation func(int, T, R) R) R {
    accumulator := initial
    if operation == nil { return accumulator }
    for i, elem := range slice {
        accumulator = operation(i, elem, accumulator)
    }
    return accumulator
}

// Accumulates value starting with the given initial 
// by performing operation on the slice from right to left
func FoldRight[T any, R any] (slice []T, initial R, operation func(T, R) R) R {
    accumulator := initial
    if operation == nil { return accumulator }
    for i := len(slice) - 1; i >= 0; i-- {
        accumulator = operation(slice[i], accumulator)
    }
    return accumulator
}


// Accumulates value starting with the given initial 
// by performing operation on the slice from right to left 
// The operation also considers index of each element
func FoldRightIndexed[T any, R any] (slice []T, initial R, operation func(int, T, R) R) R {
    accumulator := initial
    if operation == nil { return accumulator }
    for i := len(slice) - 1; i >= 0; i-- {
        accumulator = operation(i, slice[i], accumulator)
    }
    return accumulator
}

// Performs the given operation for each of element in the slice
func ForEach[T any] (slice []T, operation func(T)) {
    if operation == nil { return }
    for _, elem := range slice {
        operation(elem)
    }
}

// Performs the given operation for each of element in the slice
// The operation also considers index of each element
func ForEachIndexed[T any] (slice []T, operation func(int, T)) {
    if operation == nil { return }
    for i, elem := range slice {
        operation(i, elem)
    }
}

func GroupBy[T any, K comparable] (slice []T, selector func(T) K) map[K][]T {
    groups := make(map[K][]T)
    if selector == nil { return groups }
    for _, elem := range slice {
        groups[selector(elem)] = append(groups[selector(elem)], elem)
    }
    return groups
}

// Returns the index of the target element inside the slice.
// Returns -1 if that element does not exist.
func IndexOf[T comparable] (slice []T, target T) int {
    for i, elem := range slice {
        if elem == target {
            return i
        }
    }
    return -1
}

// Returns maximum of the slice, error if empty
func MaxOf[T constraints.Ordered] (slice []T) (T, error) {
    if (len(slice) == 0) {
        return zero[T](), errors.New("Slice is empty")
    }
    maxElem := slice[0]
    for _, elem := range slice {
        maxElem = max(maxElem, elem)
    }
    return maxElem, nil
}

// Returns maximum of slice using the comparer function.
// If a > b then comparer(a, b) > 0
func MaxOfBy[T any] (slice []T, comparer func(T, T) int) (T, error) {
    if (len(slice) == 0 || comparer == nil) {
        return zero[T](), errors.New("Either slice is empty or comparer is nil")
    }
    maxElem := slice[0]
    for _, elem := range slice {
        if comparer(elem, maxElem) > 0 {
            maxElem = elem
        }
    }
    return maxElem, nil
}

// Returns minimum of the slice, error if empty
func MinOf[T constraints.Ordered] (slice []T) (T, error) {
    if (len(slice) == 0) {
        return zero[T](), errors.New("Slice is empty")
    }
    maxElem := slice[0]
    for _, elem := range slice {
        maxElem = min(maxElem, elem)
    }
    return maxElem, nil
}

// Returns minimum of slice using the comparer function.
// If a > b then comparer(a, b) > 0
func MinOfBy[T any] (slice []T, comparer func(T, T) int) (T, error) {
    if (len(slice) == 0 || comparer == nil) {
        return zero[T](), errors.New("Either slice is empty or comparer is nil")
    }
    maxElem := slice[0]
    for _, elem := range slice {
        if comparer(elem, maxElem) <= 0 {
            maxElem = elem
        }
    }
    return maxElem, nil
}

// Partitions the slice based on the predicate.
// The left slice contains the elements that satisfies predicate 
// and the right one contains the elements that does not.
func Partition[T any] (slice []T, predicate func(T) bool) ([]T, []T) {
    left := []T{}
    right := []T{}
    if predicate == nil { return left, right }

    for _, elem := range slice {
        if predicate(elem) {
            left = append(left, elem)
        } else {
            right = append(right, elem)
        }
    }
    return left, right
}

// Returns the elements in reversed oreder
func Reversed[T any] (slice []T) []T {
    reversed := []T{}
    for i := len(slice) - 1; i >= 0; i-- {
        reversed = append(reversed, slice[i])
    }
    return reversed
} 

// Returns subarray of the slice from `from` upto `to` indecies.
// Returns an empty slice if the indecies are invalid.
func SubList[T any] (slice []T, from, to int) []T {
    sublist := []T{}
    if from < 0 || from >= len(slice) || to < 0 || to >= len(slice) {
        return sublist
    }
    for i := from; i <= to; i++ {
        sublist = append(sublist, slice[i])
    }
    return sublist
}

// Returns a slice of pointers of Pair zipping given slices.
func Zip[T, R any] (a []T, b []R) []*Pair[T, R] {
    zip := []*Pair[T, R]{}
    for i := 0; i < min(len(a), len(b)); i++ {
        zip = append(zip, &Pair[T, R]{a[i], b[i]})
    }
    return zip
}

func max[T constraints.Ordered] (a, b T) T {
    if a > b {
        return a
    }
    return b
}


func min[T constraints.Ordered] (a, b T) T {
    if a < b {
        return a
    }
    return b
}

func zero[T any]() T {
    var v T
    return v
}
