package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it/islices"
)

func TestEnumerate(t *testing.T) {
	r := Enumerate(Range(4))
	s := Collect2(r)
	assert.Equal(t, []Pair[int, int]{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, s)
}

func TestEnumerateByStep(t *testing.T) {
	r := EnumerateByStep(Range(4), 1, 2)
	s := Collect2(r)
	assert.Equal(t, []Pair[int, int]{{1, 0}, {3, 1}, {5, 2}, {7, 3}}, s)

	r = Limit2(EnumerateByStep(Range(4), 1, 2), 1)
	s = Collect2(r)
	assert.Equal(t, []Pair[int, int]{{1, 0}}, s)
}

func TestCycle(t *testing.T) {
	r := Cycle(RangeByStep(1, 3, 1))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	r = Limit(Cycle(islices.Values([]int{1, 2})), 1)
	s := islices.Collect(r)
	assert.Equal(t, []int{1}, s)

	r = Limit(Cycle(islices.Values([]int{1, 2})), 3)
	s = islices.Collect(r)
	assert.Equal(t, []int{1, 2, 1}, s)
}

func TestRepeat(t *testing.T) {
	r := Repeat(1)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	r = Limit(Repeat(1), 2)
	s := islices.Collect(r)
	assert.Equal(t, []int{1, 1}, s)
}

func TestRepeatN(t *testing.T) {
	r := RepeatN(1, 2)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit(RepeatN(1, 100), 2)
	s := islices.Collect(r)
	assert.Equal(t, []int{1, 1}, s)
}

func TestAccumulate(t *testing.T) {
	r := Accumulate(Range(4))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 6, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit(Accumulate(Range(4)), 2)
	s := islices.Collect(r)
	assert.Equal(t, []int{0, 1}, s)
}

func TestLimit(t *testing.T) {
	r := Limit(Range(4), 2)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit(Limit(Range(4), 2), 1)
	s := islices.Collect(r)
	assert.Equal(t, []int{0}, s)
}

func TestLimit2(t *testing.T) {
	r := Limit2(Enumerate(Range(4)), 2)
	next, _ := iter.Pull2(r)
	i, v, ok := next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 1, i)
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Limit2(Limit2(Enumerate(Range(4)), 2), 1)
	s := Collect2(r)
	assert.Equal(t, []Pair[int, int]{{0, 0}}, s)
}

func TestSkip(t *testing.T) {
	r := Skip(Range(4), 2)
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = Skip(Skip(Range(4), 2), 1)
	s := islices.Collect(r)
	assert.Equal(t, []int{3}, s)

	r = Limit(Skip(Range(4), 2), 1)
	s = islices.Collect(r)
	assert.Equal(t, []int{2}, s)
}

func TestZip(t *testing.T) {
	r := Zip(Range(4), Range(4))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, Zipped[int, int]{0, true, 0, true}, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, Zipped[int, int]{1, true, 1, true}, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, Zipped[int, int]{2, true, 2, true}, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, Zipped[int, int]{3, true, 3, true}, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, Zipped[int, int]{0, false, 0, false}, v)
	assert.False(t, ok)
}

func TestMap(t *testing.T) {
	r := Map(func(x int) int { return x * 2 }, Range(4))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 4, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 6, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestFilter(t *testing.T) {
	r := Filter(func(x int) bool { return x%2 == 0 }, Range(4))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}

func TestEqual(t *testing.T) {
	assert.True(t, Equal(Range(4), Range(4)))
	assert.False(t, Equal(Range(4), Range(5)))
	assert.False(t, Equal(Range(5), Range(4)))
	assert.False(t, Equal(Range(4), RangeByStep(4, 5, 1)))
	assert.False(t, Equal(RangeByStep(4, 5, 1), Range(4)))
	assert.True(t, Equal(RangeByStep(4, 5, 1), Filter(func(v int) bool {
		return true
	}, RangeByStep(4, 5, 1))))
}

func TestConcat(t *testing.T) {
	r := Concat(Range(2), Range(2))
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}
