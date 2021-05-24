package main

import "fmt"

type IQueue interface {
	Enqueue(x int)
	Dequeue() int
	Peek() int
	IsEmpty() bool
}

type Node struct {
	value int
	next *Node
}

type Queue struct {
	front *Node
	back *Node
}

func (q *Queue) IsEmpty() bool {
	if q.front == nil {
		return true
	}
	return false
}

func (q *Queue) Peek() int {
	if q.front == nil {
		return -1
	}
	return q.front.value
}

func (q *Queue) Enqueue(x int) {
	newNode := &Node{
		value: x,
	}
	if q.front == nil {
		q.front = newNode
	} else {
		q.back.next = newNode
	}
	q.back = newNode
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		return -1
	}
	value := q.front.value
	q.front = q.front.next
	return value
}

func NewQueue() IQueue {
	return &Queue{}
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
