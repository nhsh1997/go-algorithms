package main

import "fmt"

type Node struct {
	value int
	left, right *Node
}

type BST struct {
	root *Node
}

func Insert(x int, node *Node) *Node {
	if node == nil {
		newNode := &Node{value:x}
		node = newNode
	} else if x < node.value {
		node.left = Insert(x, node.left)
	} else {
		node.right = Insert(x, node.right)
	}
	return node
}

func NewBST(x int) *BST {
	return &BST{root: &Node{
		value: x,
		left:  nil,
		right: nil,
	}}
}

func main()  {
	bst := NewBST(15)
	Insert(2, bst.root)
	Insert(1, bst.root)
	Insert(19, bst.root)
	fmt.Println(bst.root.right.value)
}