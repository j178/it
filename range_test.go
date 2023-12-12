package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it/islices"
)

func TestRange(t *testing.T) {
	r := Range(3)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit(Range(3), 1)
	s := islices.Collect(r)
	assert.Equal(t, []int{0}, s)
}

func TestRangeFrom(t *testing.T) {
	r := RangeFrom(2, 3)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 5, v)
	assert.True(t, ok)

	v, ok = next()
	assert.Equal(t, 8, v)
	assert.True(t, ok)

	r = Limit(RangeFrom(2, 1), 3)
	s := islices.Collect(r)
	assert.Equal(t, []int{2, 3, 4}, s)
}

func TestRangeByStep(t *testing.T) {
	r := RangeByStep(1, 5, 2)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = RangeByStep(3, 0, -1)
	next, _ = iter.Pull(r)
	v, ok = next()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit(RangeByStep(1, 10, 2), 2)
	s := islices.Collect(r)
	assert.Equal(t, []int{1, 3}, s)
}
