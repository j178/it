package it

import (
	"iter"

	"golang.org/x/exp/constraints"
)

type Addable interface {
	constraints.Integer | constraints.Float
}

// Enumerate returns an iterator that returns (0, seq[0]), (1, seq[1]), (2, seq[2]), ...
func Enumerate[V any](seq iter.Seq[V]) iter.Seq2[int, V] {
	return EnumerateByStep(seq, 0, 1)
}

// EnumerateByStep returns an iterator that returns (start, seq[0]), (start+step, seq[1]), (start+2*step, seq[2]), ...
func EnumerateByStep[V any](seq iter.Seq[V], start, step int) iter.Seq2[int, V] {
	i := start
	return func(yield func(int, V) bool) {
		for v := range seq {
			if !yield(i, v) {
				i += step
				return
			}
			i += step
		}
	}
}

// Cycle returns an iterator returning elements from the iterable and saving a copy of each.
// When the iterable is exhausted, return elements from the saved copy. Repeats indefinitely.
func Cycle[Elem any](seq iter.Seq[Elem]) iter.Seq[Elem] {
	var (
		saved []Elem
		i     int
	)
	return func(yield func(Elem) bool) {
		for e := range seq {
			saved = append(saved, e)
			if !yield(e) {
				return
			}
		}
		for len(saved) > 0 {
			for ; i < len(saved); i++ {
				if !yield(saved[i]) {
					return
				}
			}
			i = 0
		}
	}
}

// Repeat returns an iterator that returns object indefinitely.
func Repeat[Elem any](e Elem) iter.Seq[Elem] {
	return func(yield func(Elem) bool) {
		for {
			if !yield(e) {
				break
			}
		}
	}
}

// RepeatN returns an iterator that returns object for the specified number of times.
func RepeatN[Elem any](e Elem, n int) iter.Seq[Elem] {
	i := 0
	return func(yield func(Elem) bool) {
		for ; i < n; i++ {
			if !yield(e) {
				i++
				break
			}
		}
	}
}

// Accumulate returns an iterator that returns accumulated sums.
func Accumulate[Elem Addable](seq iter.Seq[Elem]) iter.Seq[Elem] {
	var sum Elem
	return func(yield func(Elem) bool) {
		for e := range seq {
			sum += e
			if !yield(sum) {
				break
			}
		}
	}
}

// TODO AccumulateFunc
// TODO TakeWhile
// TODO DropWhile
// TODO Compress
// TODO GroupBy
// TODO FilterFalse
// TODO Starmap
// TODO Tee
// TODO Product
// TODO Permutations
// TODO Combinations
// TODO CombinationsWithReplacement
// TODO Chain
// TODO ZipLongest

// Limit returns an iterator over seq that stops after n values.
func Limit[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		if n <= 0 {
			return
		}
		for v := range seq {
			if !yield(v) {
				n--
				return
			}
			if n--; n <= 0 {
				break
			}
		}
		return
	}
}

// Limit2 returns an iterator over seq that stops after n values.
func Limit2[K any, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if n <= 0 {
			return
		}
		for k, v := range seq {
			if !yield(k, v) {
				n--
				return
			}
			if n--; n <= 0 {
				break
			}
		}
		return
	}
}

// Skip returns an iterator over seq that skips the first n values.
func Skip[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if n <= 0 {
				if !yield(v) {
					return
				}
			} else {
				n--
			}
		}
		return
	}
}

// A Zipped is a pair of zipped values, one of which may be missing,
// drawn from two different sequences.
type Zipped[V1, V2 any] struct {
	V1  V1
	Ok1 bool // whether V1 is present (if not, it will be zero)
	V2  V2
	Ok2 bool // whether V2 is present (if not, it will be zero)
}

// Zip returns an iterator that iterates x and y in parallel,
// yielding Zipped values of successive elements of x and y.
// If one sequence ends before the other, the iteration continues
// with Zipped values in which either Ok1 or Ok2 is false,
// depending on which sequence ended first.
//
// Zip is a useful building block for adapters that process
// pairs of sequences. For example, Equal can be defined as:
//
//	func Equal[V comparable](x, y Seq[V]) bool {
//		for z := range Zip(x, y) {
//			if z.Ok1 != z.Ok2 || z.V1 != z.V2 {
//				return false
//			}
//		}
//		return true
//	}
func Zip[V1, V2 any](x iter.Seq[V1], y iter.Seq[V2]) iter.Seq[Zipped[V1, V2]] {
	return func(yield func(z Zipped[V1, V2]) bool) {
		next, stop := iter.Pull(y)
		defer stop()
		v2, ok2 := next()
		for v1 := range x {
			if !yield(Zipped[V1, V2]{v1, true, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
		var zv1 V1
		for ok2 {
			if !yield(Zipped[V1, V2]{zv1, false, v2, ok2}) {
				return
			}
			v2, ok2 = next()
		}
	}
}

// Map returns an iterator over f applied to seq.
func Map[In, Out any](f func(In) Out, seq iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
		return
	}
}

// Filter returns an iterator over seq that only includes
// the values v for which f(v) is true.
func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
		return
	}
}

// Equal reports whether the two sequences are equal.
func Equal[V comparable](x, y iter.Seq[V]) bool {
	for z := range Zip(x, y) {
		if z.Ok1 != z.Ok2 || z.V1 != z.V2 {
			return false
		}
	}
	return true
}

// Concat returns an iterator over the concatenation of the sequences.
func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
		return
	}
}
