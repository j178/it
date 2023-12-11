//go:build goexperiment.rangefunc

package it

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
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
