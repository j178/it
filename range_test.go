//go:build goexperiment.rangefunc

package it

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	r := Range(3)
	v, ok := Next(r)
	assert.Equal(t, 0, v)
	assert.True(t, ok)

	v, ok = Next(r)
	assert.Equal(t, 1, v)
	assert.True(t, ok)

	v, ok = Next(r)
	assert.Equal(t, 2, v)
	assert.True(t, ok)

	v, ok = Next(r)
	assert.Equal(t, 0, v)
	assert.False(t, ok)
}
