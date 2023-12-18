package it

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// Range returns a sequence of integers from 0 (inclusive) to end (exclusive).
func Range(end int) iter.Seq[int] {
	return RangeByStep(0, end, 1)
}

// RangeByStep returns a sequence of integers from start (inclusive) to end (exclusive) by step.
// If step is zero, the sequence will be empty.
func RangeByStep[T constraints.Integer | constraints.Float](start, end, step T) iter.Seq[T] {
	i := start
	return func(yield func(v T) bool) {
		for ; (step > 0 && i < end) || (step < 0 && i > end); i += step {
			if !yield(i) {
				i += step // proceed anyway
				break
			}
		}
	}
}

// RangeFrom returns a sequence of integers from start (inclusive) to infinity by step.
// If step is zero, the sequence will be empty.
func RangeFrom[T constraints.Integer | constraints.Float](start, step T) iter.Seq[T] {
	i := start
	return func(yield func(v T) bool) {
		if step == 0 {
			return
		}
		for {
			if !yield(i) {
				i += step // proceed anyway
				break
			}
			i += step
		}
	}
}
