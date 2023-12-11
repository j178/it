//go:build goexperiment.rangefunc

package islices

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/j178/it"
)

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
