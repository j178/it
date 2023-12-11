//go:build goexperiment.rangefunc

package it

import "iter"

type Pair[K any, V any] struct {
	First  K
	Second V
}

// CollectToPairs collects key-value pairs from seq into a new slice.
func CollectToPairs[K any, V any](seq iter.Seq2[K, V]) []Pair[K, V] {
	var result []Pair[K, V]
	for k, v := range seq {
		result = append(result, Pair[K, V]{k, v})
	}
	return result
}
