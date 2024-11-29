package strategies

import "design-patterns/types"

type SDivisible[T types.CustomInt] struct {
}

func (sd *SDivisible[T]) IsOk(num T) bool {
	return num%9 == 0
}
