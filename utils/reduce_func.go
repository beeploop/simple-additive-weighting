package utils

import "iter"

func Reduce[T any](i iter.Seq[T], fn func(T, T) T, acc T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range i {
			acc = fn(acc, e)
		}

		if !yield(acc) {
			return
		}
	}
}
