package main


type Node struct {
	value *GraphVertex
	next *Node
}

type AdjacencyList struct {
	head *Node
}

type GraphVertex struct {
	name string
	edges *AdjacencyList
}
type Graph struct {
	vertices map[string]*GraphVertex
}
