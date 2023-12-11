//go:build goexperiment.rangefunc

package it

import "iter"

// Collect collects values from seq into a new slice.
func Collect[T any](seq iter.Seq[T]) []T {
	var result []T
	for v := range seq {
		result = append(result, v)
	}
	return result
}

type Pair[K any, V any] struct {
	Key K
	Val V
}

// Collect2 collects key-value pairs from seq into a new slice.
func Collect2[K any, V any](seq iter.Seq2[K, V]) []Pair[K, V] {
	var result []Pair[K, V]
	for k, v := range seq {
		result = append(result, Pair[K, V]{k, v})
	}
	return result
}
