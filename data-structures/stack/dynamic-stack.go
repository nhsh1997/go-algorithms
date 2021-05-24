package main

type Node struct {
	value int
	next *Node
}
type DynamicStack struct {
	head *Node
	count int
}

func (d *DynamicStack) push( x int) {
	newNode := &Node{
		value: x,
		next:  d.head,
	}
	d.head = newNode
	d.count ++
}

func (d *DynamicStack) pop() int {
	value := d.head.value
	d.head = d.head.next
	d.count--
	return value
}

func (d *DynamicStack) top() int {
	return d.head.value
}

func (d *DynamicStack) isEmpty() bool {
	return d.head == nil
}

func newDynamicStack() IStack {
	return &DynamicStack{}
}
