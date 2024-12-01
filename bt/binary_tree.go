package bt

import (
	"errors"
	"fmt"
)

func (bt *btr) Add(value int) {
	if bt.root == nil {
		bt.root = &node{
			value: value,
		}
		bt.len = 1
		return
	}

	depth, _ := findNodeAndAdd(value, 0, bt.root)
	if depth > bt.len {
		bt.len = depth
	}
}

func findNodeAndAdd(value int, depth int, currentNode *node) (int, *node) {
	// * let
	if value <= currentNode.value {
		fmt.Println("value left : : ", depth, value)
		if currentNode.left != nil {
			return findNodeAndAdd(value, depth+1, currentNode.left)
		} else {
			newNode := &node{
				value: value,
			}
			currentNode.left = newNode
			return depth + 1, newNode
		}
		// * gt
	} else {
		if currentNode.right != nil {
			return findNodeAndAdd(value, depth+1, currentNode.right)
		} else {
			newNode := &node{
				value: value,
			}
			currentNode.right = newNode
			return depth + 1, newNode
		}
	}
}

func (bt *btr) DisplayLeft() error {
	if bt.root == nil {
		return errors.New("no node found")
	}
	display(bt.root)
	return nil
}

func display(n *node) {
	fmt.Println("-> value : ", n.value)
	if n.left != nil {
		display(n.left)
	}
}

func (bt *btr) WholeDisplay() error {
	if bt.root == nil {
		return errors.New("no node found")
	}
	nodeList := []*node{bt.root}

	for level := 0; level <= bt.len; level++ {
		displayLine(nodeList)
		newNodeList := []*node{}
		for _, node := range nodeList {
			if node == nil {
				newNodeList = append(newNodeList, nil, nil)
			} else {
				newNodeList = append(newNodeList, node.left, node.right)
			}
		}
		nodeList = newNodeList
	}
	return nil
}

func displayLine(nodeList []*node) {
	for _, node := range nodeList {
		if node == nil {
			fmt.Printf("x\t")
		} else {
			fmt.Printf("%d\t", node.value)
		}
	}
	fmt.Printf("\n\n")
}
