package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it/islices"
)

func TestRange(t *testing.T) {
	r := Range(3)
	assert.Equal(t, []int{0, 1, 2}, islices.Collect(r))

	r = Limit(Range(3), 1)
	assert.Equal(t, []int{0}, islices.Collect(r))
}

func TestRangeFrom(t *testing.T) {
	r := RangeFrom(2, 3)
	next, stop := iter.Pull(r)
	defer stop()
	v, ok := next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 5, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 8, v)
	assert.True(t, ok)

	r = RangeFrom(0, 0)
	next, stop = iter.Pull(r)
	defer stop()
	_, ok = next()
	assert.False(t, ok)

	r = Limit(RangeFrom(2, 1), 3)
	assert.Equal(t, []int{2, 3, 4}, islices.Collect(r))
}

func TestRangeByStep(t *testing.T) {
	r := RangeByStep(1, 5, 2)
	assert.Equal(t, []int{1, 3}, islices.Collect(r))

	r = RangeByStep(3, 0, -1)
	assert.Equal(t, []int{3, 2, 1}, islices.Collect(r))

	r = RangeByStep(0, 3, 0)
	next, stop := iter.Pull(r)
	defer stop()
	_, ok := next()
	assert.False(t, ok)

	r = Limit(RangeByStep(1, 10, 2), 2)
	s := islices.Collect(r)
	assert.Equal(t, []int{1, 3}, s)
}
