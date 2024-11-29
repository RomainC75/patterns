package finder

import (
	"design-patterns/math"
	"design-patterns/types"
)

type SFinder[T types.CustomInt] struct {
	Strategy IFinderSrategy[T]
}

type IFinderSrategy[T types.CustomInt] interface {
	find() []T
}

func NewFinder[T types.CustomInt](strat IFinderSrategy[T]) *SFinder[T] {
	return &SFinder[T]{
		Strategy: strat,
	}
}

func (s *SFinder[T]) RunParallel(min T, max T) []T {
	return []T{}
}

func (s *SFinder[T]) RunSerie(min T, max T) []T {
	res := []T{}
	for i := min; i < max; i++ {
		if math.IsPrime[T](i) {
			res = append(res, i)
		}
	}
	return res
}
