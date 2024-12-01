package math

import (
	"design-patterns/types"
)

func IsPrime[T types.CustomInt](num T) bool {
	if num == 0 || num == 1 {
		return false
	}
	var i T
	for i = 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
