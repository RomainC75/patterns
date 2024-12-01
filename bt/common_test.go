package bt

import (
	"fmt"
	"testing"
)

var values = []int{
	5, 3, 7, 4, 1, 10, 9, 2,
}

func TestBinaryTree(t *testing.T) {
	bt := NewBt()
	for _, v := range values {
		bt.Add(v)
	}

	bt.WholeDisplay()
	fmt.Println("length : ", bt.len)
}
