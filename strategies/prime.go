package strategies

import "design-patterns/types"

type SPrime[T types.CustomInt] struct {
}

func NewSPrime[T types.CustomInt]() *SPrime[T] {
	return &SPrime[T]{}
}

func (sd *SPrime[T]) IsOk(num T) bool {
	if num == 1 {
		return true
	}
	var i T
	for i = 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
