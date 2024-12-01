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
		bt.len = depth + 1
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

	leftMargin := bt.len - 1
	for level := 0; level < bt.len; level++ {
		displayLine(nodeList, leftMargin, level)
		newNodeList := []*node{}
		for _, node := range nodeList {
			if node == nil {
				newNodeList = append(newNodeList, nil, nil)
			} else {
				newNodeList = append(newNodeList, node.left, node.right)
			}
		}
		nodeList = newNodeList
		leftMargin--
	}
	return nil
}

func displayLine(nodeList []*node, leftMargin int, level int) {
	for i := 0; i < leftMargin; i++ {

		fmt.Print("\t")
	}
	for index, node := range nodeList {
		if node == nil {
			fmt.Printf("x")
		} else {
			fmt.Printf("%d", node.value)
		}

		if index%2 != 0 {
			fmt.Print("\t")
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Printf("\n\n")
}
