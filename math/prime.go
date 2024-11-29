package math

import (
	"design-patterns/types"
)

func IsPrime[T types.CustomInt](num T) bool {
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
