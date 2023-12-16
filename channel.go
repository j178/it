package it

import (
	"iter"
)

// FromChannel returns a sequence from a channel.
func FromChannel[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range ch {
			if !yield(v) {
				break
			}
		}
	}
}

func ToChannel[T any](seq iter.Seq[T]) <-chan T {
	ch := make(chan T)
	go func() {
		for v := range seq {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
