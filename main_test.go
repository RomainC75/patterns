package main

import (
	"design-patterns/finder"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = struct {
	max int64
	res []int64
}{
	max: 25,
	res: []int64{
		2, 3, 5, 7, 11, 13, 17, 19,
	},
}

func TestPrimeParallel(t *testing.T) {
	f, _ := finder.NewFinder[int64](finder.Prime)
	res := f.RunParallel(0, testCases.max)
	assert.Equal(t, len(testCases.res), len(res))
	for _, val := range testCases.res {
		var found bool
		for _, pn := range res {
			if pn == val {
				found = true
				break
			}
		}
		fmt.Println("----", val)
		assert.Equal(t, found, true)
	}
	fmt.Println("res : ", res)
}

var start int64 = 1_000_000
var end int64 = 2_000_000

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
