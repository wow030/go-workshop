package fib_test

import (
	"go-workshop/fib"
	"testing"
)

func BenchmarkFib20(b *testing.B) {
	// if we have boring and expensive setup
	// then reset timer
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fib.Fib(20)
	}
}
