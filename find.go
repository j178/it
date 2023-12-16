package it

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// IndexOf returns the first index of v in seq, or -1 if v is not present.
func IndexOf[T comparable](seq iter.Seq[T], v T) int {
	i := 0
	for e := range seq {
		if e == v {
			return i
		}
		i++
	}
	return -1
}

// Find searches for an element satisfying f in seq. If found, it returns the element and its index and true.
// Otherwise, it returns the zero value, -1, and false.
func Find[T any](seq iter.Seq[T], f func(T) bool) (T, int, bool) {
	i := 0
	for e := range seq {
		if f(e) {
			return e, i, true
		}
		i++
	}
	var z T
	return z, -1, false
}

// FindOrElse searches for an element satisfying f in seq. If found, it returns the element. Otherwise, it returns fallback.
func FindOrElse[T any](seq iter.Seq[T], fallback T, f func(T) bool) T {
	if e, _, ok := Find(seq, f); ok {
		return e
	}
	return fallback
}

// Contains returns true if seq contains v.
func Contains[T comparable](seq iter.Seq[T], v T) bool {
	for e := range seq {
		if e == v {
			return true
		}
	}
	return false
}

// ContainsBy returns true if seq contains an element that satisfies f.
func ContainsBy[T any](seq iter.Seq[T], f func(T) bool) bool {
	for e := range seq {
		if f(e) {
			return true
		}
	}
	return false
}

// Max returns the maximum element of seq.
func Max[T constraints.Ordered](seq iter.Seq[T]) (T, bool) {
	var (
		result T
		ok     bool
	)
	next, stop := iter.Pull(seq)
	result, ok = next()
	if !ok {
		stop()
		return result, false
	}
	for e, ok := next(); ok; e, ok = next() {
		if result < e {
			result = e
		}
	}
	stop()
	return result, true
}

// Min returns the minimum element of seq.
func Min[T constraints.Ordered](seq iter.Seq[T]) (T, bool) {
	var (
		result T
		ok     bool
	)
	next, stop := iter.Pull(seq)
	result, ok = next()
	if !ok {
		stop()
		return result, false
	}
	for e, ok := next(); ok; e, ok = next() {
		if result > e {
			result = e
		}
	}
	stop()
	return result, true
}

// MaxBy returns the maximum element of seq according to f.
func MaxBy[T any](seq iter.Seq[T], f func(T, T) bool) (T, bool) {
	var (
		result T
		ok     bool
	)
	next, stop := iter.Pull(seq)
	result, ok = next()
	if !ok {
		stop()
		return result, false
	}
	for e, ok := next(); ok; e, ok = next() {
		if f(e, result) {
			result = e
		}
	}
	stop()
	return result, true
}
