package it

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it/islices"
)

func TestFromChannel(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	r := FromChannel(ch)
	assert.Equal(t, []int{1, 2, 3}, islices.Collect(r))

	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	r = Limit(FromChannel(ch), 2)
	assert.Equal(t, []int{1, 2}, islices.Collect(r))
}

func TestToChannel(t *testing.T) {
	r := Range(3)
	ch := ToChannel(r)
	assert.Equal(t, 0, <-ch)
	assert.Equal(t, 1, <-ch)
	assert.Equal(t, 2, <-ch)
	_, ok := <-ch
	assert.False(t, ok)
}
