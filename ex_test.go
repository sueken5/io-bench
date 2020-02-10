package io_bench_test

import (
	io_bench "github.com/sueken5/io-bench"
	"testing"
)

func BenchmarkEx1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := io_bench.Ex1(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEx2(b *testing.B) {
	if err := io_bench.Ex2(); err != nil {
		b.Error(err)
	}
}
