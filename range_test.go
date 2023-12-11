//go:build goexperiment.rangefunc

package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
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
}
