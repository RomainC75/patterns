package finder

import (
	"design-patterns/strategies"
	"design-patterns/types"
	"errors"
)

type SFinder[T types.CustomInt] struct {
	Strategy IFinderSrategy[T]
}

type IFinderSrategy[T types.CustomInt] interface {
	IsOk(num T) bool
}

func NewFinder[T types.CustomInt](scenario EScenario) (*SFinder[T], error) {
	var strat IFinderSrategy[T]
	switch scenario {
	case Prime:
		strat = new(strategies.SPrime[T])
	case Divisible:
		strat = new(strategies.SDivisible[T])
	default:
		return &SFinder[T]{}, errors.New("unknown scenario")
	}

	return &SFinder[T]{
		Strategy: strat,
	}, nil
}

func (s *SFinder[T]) RunSerie(min T, max T) []T {
	res := []T{}
	for i := min; i < max; i++ {
		if s.Strategy.IsOk(i) {
			res = append(res, i)
		}
	}
	return res
}

func (s *SFinder[T]) RunParallel(min T, max T) []T {
	res := []T{}
	for i := min; i < max; i++ {
		if s.Strategy.IsOk(i) {
			res = append(res, i)
		}
	}
	return res
}
