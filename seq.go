//go:build goexperiment.rangefunc

package it

import "iter"

type Pair[K any, V any] struct {
	First  K
	Second V
}

// Collect collects key-value pairs from seq into a new slice.
func Collect[K any, V any](seq iter.Seq2[K, V]) []Pair[K, V] {
	var result []Pair[K, V]
	for k, v := range seq {
		result = append(result, Pair[K, V]{k, v})
	}
	return result
}

// TODO: figure out how this work

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
