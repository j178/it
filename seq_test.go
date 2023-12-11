//go:build goexperiment.rangefunc

package it

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	r := Range(10)
	assert.Equal(t, 10, Count(r))
	assert.Equal(t, 0, Count(r))

	r = Range(0)
	assert.Equal(t, 0, Count(r))

	r = Range(-1)
	assert.Equal(t, 0, Count(r))
}
