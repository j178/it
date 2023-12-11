//go:build goexperiment.rangefunc

package it

import "iter"

// Range returns a sequence of integers from 0 to end (exclusive).
func Range(end int) iter.Seq[int] {
	return RangeByStep(0, end, 1)
}

// RangeByStep returns a sequence of integers from start to end (exclusive) by step.
func RangeByStep(start, end, step int) iter.Seq[int] {
	i := start
	return func(yield func(v int) bool) {
		for ; (step > 0 && i < end) || (step < 0 && i > end); i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// RangeFrom returns a sequence of integers from start (inclusive) to infinity by step.
func RangeFrom(start, step int) iter.Seq[int] {
	i := start
	return func(yield func(v int) bool) {
		for ; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}
