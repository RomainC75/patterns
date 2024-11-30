package main

import (
	"design-patterns/finder"
	"fmt"
	"testing"
)

var start int64 = 1_000_000
var end int64 = 1_100_000

func BenchmarkPrimeParallel(b *testing.B) {
	f, _ := finder.NewFinder[int64](finder.Prime)
	res := f.RunParallel(start, end)
	fmt.Println("-- Parallel -->", len(res))
}
func BenchmarkPrimeSerie(b *testing.B) {
	f, _ := finder.NewFinder[int64](finder.Prime)
	res := f.RunSerie(start, end)
	fmt.Println("-- Serie -->", len(res))
}
