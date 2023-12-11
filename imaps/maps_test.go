//go:build goexperiment.rangefunc

package imaps

import (
	"cmp"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it"
	"github.com/j178/it/islices"
)

func TestAll(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	m1 := it.Collect(All(m))
	slices.SortFunc(m1, func(a, b it.Pair[int, int]) int { return cmp.Compare(a.First, b.First) })
	assert.Equal(t, []it.Pair[int, int]{{1, 1}, {2, 2}, {3, 3}}, m1)
}

func TestKeys(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	r := Keys(m)
	assert.Equal(t, []int{1, 2, 3}, islices.Sorted(r))
}

func TestValues(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	r := Values(m)
	assert.Equal(t, []int{1, 2, 3}, islices.Sorted(r))
}

func TestInsert(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	m = Insert(m, All(map[int]int{0: 0}))
	assert.Equal(t, map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, m)
}

func TestCollect(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	m1 := Collect(All(m))
	assert.Equal(t, m, m1)
}
