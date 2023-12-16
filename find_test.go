package it

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	r := Range(4)
	i := IndexOf(r, 2)
	assert.Equal(t, 2, i)

	r = Range(4)
	i = IndexOf(r, 4)
	assert.Equal(t, -1, i)
}

func TestFind(t *testing.T) {
	r := Range(4)
	_, i, ok := Find(r, func(v int) bool { return v == 2 })
	assert.Equal(t, 2, i)
	assert.True(t, ok)

	r = Range(4)
	_, i, ok = Find(r, func(v int) bool { return v == 4 })
	assert.Equal(t, -1, i)
	assert.False(t, ok)
}

func TestFindOrElse(t *testing.T) {
	r := Range(4)
	i := FindOrElse(r, -1, func(v int) bool { return v == 2 })
	assert.Equal(t, 2, i)

	r = Range(4)
	i = FindOrElse(r, 0, func(v int) bool { return v == 4 })
	assert.Equal(t, 0, i)
}

func TestContains(t *testing.T) {
	assert.True(t, Contains(Range(4), 2))
	assert.False(t, Contains(Range(4), 4))
	assert.False(t, Contains(Range(4), 5))
}

func TestContainsBy(t *testing.T) {
	assert.True(t, ContainsBy(Range(4), func(v int) bool { return v == 2 }))
	assert.False(t, ContainsBy(Range(4), func(v int) bool { return v == 4 }))
	assert.False(t, ContainsBy(Range(4), func(v int) bool { return v == 5 }))
}

func TestMax(t *testing.T) {
	r := Range(0)
	v, ok := Max(r)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Range(4)
	v, ok = Max(r)
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	r = RangeByStep(3, 0, -1)
	v, ok = Max(r)
	assert.Equal(t, 3, v)
	assert.True(t, ok)
}

func TestMin(t *testing.T) {
	r := Range(0)
	v, ok := Min(r)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Range(4)
	v, ok = Min(r)
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	r = RangeByStep(3, 0, -1)
	v, ok = Min(r)
	assert.Equal(t, 1, v)
	assert.True(t, ok)
}

func TestMaxBy(t *testing.T) {
	r := Range(0)
	v, ok := MaxBy(r, func(a, b int) bool { return a > b })
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Range(4)
	v, ok = MaxBy(r, func(a, b int) bool { return a > b })
	assert.Equal(t, 3, v)
	assert.True(t, ok)

	r = RangeByStep(3, 0, -1)
	v, ok = MaxBy(r, func(a, b int) bool { return a > b })
	assert.Equal(t, 3, v)
	assert.True(t, ok)
}
