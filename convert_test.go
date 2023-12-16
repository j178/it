package it

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it/islices"
)

func TestFirsts(t *testing.T) {
	r := islices.All([]int{1, 1, 1})
	r1 := Firsts(r)
	assert.Equal(t, []int{0, 1, 2}, islices.Collect(r1))

	r2 := Limit(Firsts(r), 2)
	assert.Equal(t, []int{0, 1}, islices.Collect(r2))
}

func TestSeconds(t *testing.T) {
	r := islices.All([]int{1, 2, 3})
	r1 := Seconds(r)
	assert.Equal(t, []int{1, 2, 3}, islices.Collect(r1))

	r2 := Limit(Seconds(r), 2)
	assert.Equal(t, []int{1, 2}, islices.Collect(r2))
}
