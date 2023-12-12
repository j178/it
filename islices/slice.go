package islices

import (
	"cmp"
	"iter"
	"slices"
)

// All returns an iterator over index-value pairs in the slice.
// The indexes range in the usual order, from 0 through len(s)-1.
func All[Slice ~[]Elem, Elem any](s Slice) iter.Seq2[int, Elem] {
	return func(yield func(int, Elem) bool) {
		for i, v := range s {
			if !yield(i, v) {
				break
			}
		}
	}
}

// Backward returns an iterator over index-value pairs in the slice,
// traversing it backward. The indexes range from len(s)-1 down to 0.
func Backward[Slice ~[]Elem, Elem any](s Slice) iter.Seq2[int, Elem] {
	return func(yield func(int, Elem) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				break
			}
		}
	}
}

// Values returns an iterator over the values in the slice,
// starting with s[0].
func Values[Slice ~[]Elem, Elem any](s Slice) iter.Seq[Elem] {
	return func(yield func(Elem) bool) {
		for _, v := range s {
			if !yield(v) {
				break
			}
		}
	}
}

// Append appends the values from seq to the slice and returns the extended slice.
func Append[Slice ~[]Elem, Elem any](x Slice, seq iter.Seq[Elem]) Slice {
	for v := range seq {
		x = append(x, v)
	}
	return x
}

// Collect collects values from seq into a new slice and returns it.
func Collect[Elem any](seq iter.Seq[Elem]) []Elem {
	return Append([]Elem(nil), seq)
}

// Sorted collects values from seq into a new slice, sorts the slice, and returns it.
func Sorted[Elem cmp.Ordered](seq iter.Seq[Elem]) []Elem {
	slice := Collect(seq)
	slices.Sort(slice)
	return slice
}

// SortedFunc collects values from seq into a new slice, sorts the slice, and returns it.
func SortedFunc[Elem any](seq iter.Seq[Elem], cmp func(Elem, Elem) int) []Elem {
	slice := Collect(seq)
	slices.SortFunc(slice, cmp)
	return slice
}
