package imaps

import "iter"

// All returns an iterator over key-value pairs from m.
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

// Keys returns an iterator over keys in m.
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range m {
			if !yield(k) {
				break
			}
		}
	}
}

// Values returns an iterator over values in m.
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				break
			}
		}
	}
}

// Insert adds the key-value pairs from seq to m.
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V]) Map {
	for k, v := range seq {
		m[k] = v
	}
	return m
}

// Collect collects key-value pairs from seq into a new map
// and returns it.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	m := make(map[K]V)
	Insert(m, seq)
	return m
}
