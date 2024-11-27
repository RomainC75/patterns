package main

import "testing"

var start int64 = 1_000_000
var end int64 = 2_000_000

func BenchmarkParallel(b *testing.B) {
	parallel(start, end)
}

func BenchmarkSerie(b *testing.B) {
	serie(start, end)
}
