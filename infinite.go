package it

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type Addable interface {
	constraints.Integer | constraints.Float
}

// Count returns an iterator that returns evenly spaced values starting with number start.
func Count[I Addable](start, step I) iter.Seq[I] {
	return func(yield func(I) bool) {
		for i := start; ; i += step {
			if !yield(i) {
				break
			}
		}
	}
}

// Cycle returns an iterator returning elements from the iterable and saving a copy of each.
// When the iterable is exhausted, return elements from the saved copy. Repeats indefinitely.
func Cycle[Elem any](iterable iter.Seq[Elem]) iter.Seq[Elem] {
	return func(yield func(Elem) bool) {
		var saved []Elem
		for e := range iterable {
			if !yield(e) {
				return
			}
			saved = append(saved, e)
		}
		for len(saved) > 0 {
			for _, e := range saved {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// Repeat returns an iterator that returns object for the specified number of times.
// If times < 0, Repeat runs indefinitely.
func Repeat[Elem any](e Elem, times int) iter.Seq[Elem] {
	return func(yield func(Elem) bool) {
		for i := 0; times < 0 || i < times; i++ {
			if !yield(e) {
				break
			}
		}
	}
}
