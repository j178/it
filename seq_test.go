//go:build goexperiment.rangefunc

package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func R(stop int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < stop; i++ {
			if !yield(i, i) {
				return
			}
		}
	}
}

func TestCollect(t *testing.T) {
	r := R(4)
	assert.Equal(t, []Pair[int, int]{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, Collect(r))
}

func TestNext(t *testing.T) {
	r := Range(4)
	v, ok := Next(r)
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	r = RangeFrom(2, 1)
	v, ok = Next(r)
	assert.Equal(t, 2, v)
	assert.True(t, ok)
}

func TestNth(t *testing.T) {
	r := Range(4)
	v, ok := Nth(r, 2)
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	r = RangeFrom(2, 1)
	v, ok = Nth(r, 2)
	assert.Equal(t, 4, v)
	assert.True(t, ok)
}

func TestLast(t *testing.T) {
	r := Range(4)
	v, ok := Last(r)
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	r = Range(0)
	v, ok = Last(r)
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestCount(t *testing.T) {
	r := Range(10)
	assert.Equal(t, 10, Count(r))
	assert.Equal(t, 0, Count(r))

	r = Range(0)
	assert.Equal(t, 0, Count(r))

	r = Range(-1)
	assert.Equal(t, 0, Count(r))
}
