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
	m1 := it.Collect2(All(m))
	slices.SortFunc(m1, func(a, b it.Pair[int, int]) int { return cmp.Compare(a.First, b.First) })
	assert.Equal(t, []it.Pair[int, int]{{1, 1}, {2, 2}, {3, 3}}, m1)

	i := 0
	for _, _ = range All(m) {
		if i >= 1 {
			break
		}
		i++
	}
}

func TestKeys(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	r := Keys(m)
	assert.Equal(t, []int{1, 2, 3}, islices.Sorted(r))

	i := 0
	for _ = range Keys(m) {
		if i >= 1 {
			break
		}
		i++
	}
}

func TestValues(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	r := Values(m)
	assert.Equal(t, []int{1, 2, 3}, islices.Sorted(r))

	i := 0
	for _ = range Values(m) {
		if i >= 1 {
			break
		}
		i++
	}
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
