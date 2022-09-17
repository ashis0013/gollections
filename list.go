package gollections

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Filter[T any] (slice []T, predicate func(T) bool) []T {
    var filteredSlice = []T{}
    for _, elem := range slice {
        if predicate(elem) {
            filteredSlice = append(filteredSlice, elem)
        }
    }
    return filteredSlice
}

func FilterIndexed[T any] (slice []T, predicate func(int, T) bool) []T {
    var filteredSlice = []T{}
    for i, elem := range slice {
        if predicate(i, elem) {
            filteredSlice = append(filteredSlice, elem)
        }
    }
    return filteredSlice
}

func Fmap[T any, R any] (slice []T, transform func(T) R) []R {
    var transformedSlice = []R{}
    for _, elem := range slice {
        transformedSlice = append(transformedSlice, transform(elem))
    }
    return transformedSlice 
}

func All[T any] (slice []T, predicate func(T) bool) bool {
    for _, elem := range slice {
        if !predicate(elem) {
            return false
        }
    }
    return true
} 

func Any[T any] (slice []T, predicate func(T) bool) bool {
    for _, elem := range slice {
        if predicate(elem) {
            return true
        }
    }
    return false
}

func Associate[T any, K comparable, V any] (slice []T, transform func(T) (K, V)) map[K]V {
    dict := make(map[K]V)
    for _, elem := range slice {
        key, value := transform(elem)
        dict[key] = value
    }
    return dict
}

func Contains[T comparable] (slice []T, target T) bool {
    for _, elem := range slice {
        if elem == target {
            return true
        }
    }
    return false
}

func Count[T any] (slice []T, predicate func(T) bool) int {
    count := 0
    for _, elem := range slice {
        if predicate(elem) {
            count++
        }
    }
    return count
}

func Drop[T any] (slice []T, index int) []T {
    if index < 0 || index > len(slice) {
        return slice
    }
    return FilterIndexed(slice, func(i int, elem T) bool {
        return i != index
    })
}

func First[T any] (slice []T, predicate func(T) bool) (T, error) {
    for _, elem := range slice {
        if predicate(elem) {
            return elem, nil
        }
    }
    var zero T
    return zero, errors.New("No element found satisfying predicate")
}

func FirstOrDefault[T any] (slice []T, defaultValue T, predicate func(T) bool) T {
    first, err := First(slice, predicate)
    if err != nil {
        return defaultValue
    }
    return first
}

func Flatten[T any] (slices [][]T) []T {
    var flattenedSlice []T
    for _, slice := range slices {
        for _, elem := range slice {
            flattenedSlice = append(flattenedSlice, elem)
        }
    }
    return flattenedSlice
}

func Fold[T any, R any] (slice []T, initial R, operation func(T, R) R) R {
    accumulator := initial
    for _, elem := range slice {
        accumulator = operation(elem, accumulator)
    }
    return accumulator
}

func FoldIndexed[T any, R any] (slice []T, initial R, operation func(int, T, R) R) R {
    accumulator := initial
    for i, elem := range slice {
        accumulator = operation(i, elem, accumulator)
    }
    return accumulator
}

func FoldRight[T any, R any] (slice []T, initial R, operation func(T, R) R) R {
    accumulator := initial
    for i := len(slice) - 1; i >= 0; i-- {
        accumulator = operation(slice[i], accumulator)
    }
    return accumulator
}


func FoldRightIndexed[T any, R any] (slice []T, initial R, operation func(int, T, R) R) R {
    accumulator := initial
    for i := len(slice) - 1; i >= 0; i-- {
        accumulator = operation(i, slice[i], accumulator)
    }
    return accumulator
}

// sus
func GroupBy[T any, K comparable] (slice []T, selector func(T) K) map[K][]T {
    groups := make(map[K][]T)
    for _, elem := range slice {
        groups[selector(elem)] = append(groups[selector(elem)], elem)
    }
    return groups
}

func IndexOf[T comparable] (slice []T, target T) int {
    for i, elem := range slice {
        if elem == target {
            return i
        }
    }
    return -1
}

func MaxOf[T constraints.Ordered] (slice []T) (T, error) {
    if (len(slice) == 0) {
        var zero T
        return zero, nil
    }
    maxElem := slice[0]
    for _, elem := range slice {
        maxElem = max(maxElem, elem)
    }
    return maxElem, nil
}

func MaxOfBy[T any] (slice []T, comparer func(T, T) int) (T, error) {
    if (len(slice) == 0) {
        var zero T
        return zero, nil
    }
    maxElem := slice[0]
    for _, elem := range slice {
        if comparer(elem, maxElem) > 0 {
            maxElem = elem
        }
    }
    return maxElem, nil
}


func MinOf[T constraints.Ordered] (slice []T) (T, error) {
    if (len(slice) == 0) {
        var zero T
        return zero, nil
    }
    maxElem := slice[0]
    for _, elem := range slice {
        maxElem = min(maxElem, elem)
    }
    return maxElem, nil
}

func MinOfBy[T any] (slice []T, comparer func(T, T) int) (T, error) {
    if (len(slice) == 0) {
        var zero T
        return zero, nil
    }
    maxElem := slice[0]
    for _, elem := range slice {
        if comparer(elem, maxElem) <= 0 {
            maxElem = elem
        }
    }
    return maxElem, nil
}

func Partition[T any] (slice []T, predicate func(T) bool) ([]T, []T) {
    left :=  []T{}
    right :=  []T{}

    for _, elem := range slice {
        if predicate(elem) {
            left = append(left, elem)
        } else {
            right = append(right, elem)
        }
    }
    return left, right
}

func Reversed[T any] (slice []T) []T {
    reversed := []T{}
    for i := len(slice) - 1; i <= 0; i-- {
        reversed = append(reversed, slice[i])
    }
    return reversed
} 


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

// sus
func Zip[T any] (a []T, b []T) [][]T {
    lenght := min(len(a), len(b))
    zip := make([][]T, lenght)
    for i := 0; i <= lenght; i++ {
        zip = append(zip, []T{a[i], b[i]})
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
