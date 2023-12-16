package it

import "iter"

type Pair[K any, V any] struct {
	First  K
	Second V
}

// Collect2 collects key-value pairs from seq into a new slice.
func Collect2[K any, V any](seq iter.Seq2[K, V]) []Pair[K, V] {
	var result []Pair[K, V]
	for k, v := range seq {
		result = append(result, Pair[K, V]{k, v})
	}
	return result
}

// Next returns the first element of seq.
func Next[T any](seq iter.Seq[T]) (T, bool) {
	return Nth(seq, 0)
}

// Nth returns the nth element of seq.
func Nth[T any](seq iter.Seq[T], n int) (T, bool) {
	i := 0
	for v := range seq {
		if i == n {
			return v, true
		}
		i++
	}
	var result T
	return result, false
}

// Last returns the last element of seq.
func Last[T any](seq iter.Seq[T]) (T, bool) {
	var (
		result T
		ok     bool
	)
	for v := range seq {
		result, ok = v, true
	}
	return result, ok
}

// Count returns the number of elements in seq.
func Count[T any](seq iter.Seq[T]) (cnt int) {
	// `for range seq` cannot compile
	for _ = range seq {
		cnt++
	}
	return cnt
}

// Count2 returns the number of elements in seq.
func Count2[K any, V any](seq iter.Seq2[K, V]) (cnt int) {
	for _, _ = range seq {
		cnt++
	}
	return cnt
}

// ForEach calls f for each element of seq.
func ForEach[V any](seq iter.Seq[V], f func(V)) {
	for v := range seq {
		f(v)
	}
}

// All returns true if all elements of seq satisfy f.
func All[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if !f(v) {
			return false
		}
	}
	return true
}

// Any returns true if any element of seq satisfies f.
func Any[V any](seq iter.Seq[V], f func(V) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}
