package gollections

type Pair[T, R any] struct {
    First T
    Second R
}

func Entries[K comparable, V any] (hashMap map[K]V) []*Pair[K, V] {
    entries := []*Pair[K, V]{}
    for key, value := range hashMap {
        entries = append(entries, &Pair[K, V]{key, value})
    }
    return entries
}

func Keys[K comparable, V any] (hashMap map[K]V) []K {
    keys := []K{}
    for key := range hashMap {
        keys = append(keys, key)
    }
    return keys
}

func Values[K comparable, V any] (hashMap map[K]V) []V {
    values := []V{}
    for _, value := range hashMap {
        values = append(values, value)
    }
    return values
}

func ContainsKey[K comparable, V any] (hashMap map[K]V, target K) bool {
    _, found := hashMap[target]
    return found
}

func GetOrDefault[K comparable, V any] (hashMap map[K]V, key K, defaultVal V) V {
    if value, found := hashMap[key]; found {
        return value
    }
    return defaultVal
}

func FilterKeys[K comparable, V any] (hashMap map[K]V, predicate func(K) bool) map[K]V {
    filtered := make(map[K]V)
    for key, value := range hashMap {
        if predicate(key) {
            filtered[key] = value
        }
    }
    return filtered
}

func FlatMap[K comparable, V, R any] (hashMap map[K]V, transform func(K, V) R) []R {
    slice := []R{}
    for key, value := range hashMap {
        slice = append(slice, transform(key, value))
    }
    return slice
}

func ForEachEntry[K comparable, V any] (hashMap map[K]V, operation func(K, V)) {
    for key, value := range hashMap {
        operation(key, value)
    }
}
