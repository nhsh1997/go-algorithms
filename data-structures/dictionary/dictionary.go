package main

type IDictionary interface {
	Insert(key int, value int)
	Delete(key int, value int)
	LookUp(key int)
}

type DictNode struct {
	key int
	value int
	next *DictNode
}

