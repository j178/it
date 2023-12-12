package islices

import (
	"cmp"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it"
)

func TestAll(t *testing.T) {
	r := All([]int{1, 2, 3})
	next, _ := iter.Pull2(r)
	i, v, ok := next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 1, i)
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 2, i)
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = it.Limit2(All([]int{1, 2, 3}), 1)
	assert.Equal(t, []it.Pair[int, int]{{0, 1}}, it.Collect2(r))
}

func TestBackward(t *testing.T) {
	r := Backward([]int{1, 2, 3})
	next, _ := iter.Pull2(r)
	i, v, ok := next()
	assert.Equal(t, 2, i)
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 1, i)
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	i, v, ok = next()
	assert.Equal(t, 0, i)
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = it.Limit2(Backward([]int{1, 2, 3}), 1)
	assert.Equal(t, []it.Pair[int, int]{{2, 3}}, it.Collect2(r))
}

func TestValues(t *testing.T) {
	r := Values([]int{1, 2, 3})
	next, _ := iter.Pull(r)
	v, ok := next()
	assert.Equal(t, 1, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 2, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 3, v)
	assert.True(t, ok)
	v, ok = next()
	assert.Equal(t, 0, v)
	assert.False(t, ok)

	r = it.Limit(Values([]int{1, 2, 3}), 1)
	assert.Equal(t, []int{1}, Collect(r))
}

func TestAppend(t *testing.T) {
	r := Append([]int{1, 2, 3}, it.RangeByStep(4, 6, 1))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r)
}

func TestCollect(t *testing.T) {
	r := it.Range(10)
	s := Collect(r)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)

	s = Collect(
		func(yield func(v int) bool) {
			return
		},
	)
	assert.Equal(t, 0, len(s))

	s1 := Collect(
		func(yield func(str string) bool) {
			if !yield("hello") {
				return
			}
		},
	)
	assert.Equal(t, []string{"hello"}, s1)

	s2 := Collect(
		func(yield func(str string) bool) {
			if !yield("hello") {
				return
			}
			if !yield("world") {
				return
			}
		},
	)
	assert.Equal(t, []string{"hello", "world"}, s2)
}

func TestSorted(t *testing.T) {
	r := it.RangeByStep(3, 0, -1)
	s := Sorted(r)
	assert.Equal(t, []int{1, 2, 3}, s)
}

func TestSortedFunc(t *testing.T) {
	r := it.RangeByStep(3, 0, -1)
	s := SortedFunc(r, func(a, b int) int { return cmp.Compare(a, b) })
	assert.Equal(t, []int{1, 2, 3}, s)
}
