package fibonacci

import (
	"testing"
)

var UPTO int = 42

func BenchmarkClosedForm(b *testing.B) {
	for b.Loop() {
		for i := range UPTO {
			ClosedForm(i)
		}
	}
}

func BenchmarkRecursive(b *testing.B) {
	for b.Loop() {
		for i := range UPTO {
			Recursive(i)
		}
	}
}

func BenchmarkIterative(b *testing.B) {
	for b.Loop() {
		for i := range UPTO {
			Iterative(i)
		}
	}
}

func BenchmarkMemoized(b *testing.B) {
	for b.Loop() {
		for i := range UPTO {
			MemoizedRecurssive(i)
		}
	}
}
