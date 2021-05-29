package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type Stack struct {
	head *Node
	length int
}

func (q *Stack) Push(value int) {
	q.head = &Node{
		value: value,
		next:  q.head,
	}
	q.length++
}

func (q *Stack) Pop() int {
	if q.head == nil {
		return -1
	}
	value := q.head.value
	q.head = q.head.next
	return value
}

type IStack interface {
	Push(value int)
	Pop() int
}

func NewStack() IStack {
	return &Stack{}
}

type ILinkedList interface{
	Display()
	InsertItem(value int)
	Reverse()
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}

func (l *LinkedList) Reverse() {
	stack := NewStack()
	current := l.head
	for current != nil {
		value := current.value
		stack.Push(value)
		current = current.next
	}

	current = l.head
	for current != nil  {
		value := stack.Pop()
		current.value = value
		current = current.next
	}
}

func (l *LinkedList) InsertItem(value int) {
	newNode := &Node{
		value: value,
	}
	if l.head == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
	l.length++
}

func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func NewLinkedList() ILinkedList {
	return &LinkedList{}
}

func main() {
	linkedList := NewLinkedList()
	linkedList.InsertItem(1)
	linkedList.InsertItem(2)
	linkedList.InsertItem(3)
	linkedList.InsertItem(4)
	linkedList.InsertItem(12)
	linkedList.Display()
	linkedList.Reverse()
	linkedList.Display()
}

