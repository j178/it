package it

import (
	"iter"
)

// Firsts returns a sequence of the first elements of a sequence of pairs.
func Firsts[K any, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range seq {
			if !yield(k) {
				break
			}
		}
	}
}

// Seconds returns a sequence of the second elements of a sequence of pairs.
func Seconds[K any, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				break
			}
		}
	}
}

// TODO Seq2 to Seq
