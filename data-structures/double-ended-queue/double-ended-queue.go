package main

type Node struct {
	value int
	next *Node
	prev *Node
}

type DoubleEndedQueue struct {
	front *Node
	back *Node
}


