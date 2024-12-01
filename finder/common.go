package finder

import (
	"design-patterns/strategies"
	"design-patterns/types"
	"errors"
	"sync"
)

type SFinder[T types.CustomInt] struct {
	res      []T
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
	// s.res = make([]T, 0, (max-min)/100)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	numbers, done := s.Generator(min, max, wg)
	res := s.Parallel(8, numbers)

	s.Accumulator(1, res, done)

	wg.Wait()

	return s.res
}

func (s *SFinder[T]) Generator(min T, max T, wg *sync.WaitGroup) (chan T, chan int) {
	gen := make(chan T)
	done := make(chan int)
	go func() {
		defer close(gen)
		defer close(done)
		for i := min; i <= max; i++ {
			gen <- i
		}
		wg.Done()
	}()
	return gen, done
}

func (s *SFinder[T]) Parallel(chanNum int, inputChan chan T) chan T {
	// ! capacity ?

	// out := make([]chan T, chanNum)
	res := make(chan T)

	// for i := 0; i < chanNum; i++ {
	// 	ch := make(chan T)
	// 	out = append(out, ch)
	// }

	for i := 0; i < chanNum; i++ {
		go func() {
			for {
				j := <-inputChan
				if s.Strategy.IsOk(j) {
					res <- j
					// fmt.Println(i, " -> ", j)
				}
			}
		}()
	}
	return res
}

func (s *SFinder[T]) Accumulator(chanNum int, inputChan chan T, done chan int) {

	go func() {
		// resCh := make(chan []T)

		for {
			select {
			case <-done:
				return
			case val := <-inputChan:
				// fmt.Println("- ", val)
				s.res = append(s.res, val)
			}
			// res = append(res, val)
		}
	}()
}
