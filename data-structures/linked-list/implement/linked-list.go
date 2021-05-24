package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	length int
}


func getItem(head *Node, n int) int {
	count := 0
	current := head
	for current != nil {
		if count == n {
			return current.value
		}
		count++
		current = current.next
	}
	return -1
}

func addItem(head *Node, node *Node) {
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func search(head *Node, x int) bool {
	for head != nil {
		if head.value == x {
			return true
		}
	}
	return false
}

func insertFront(head *Node, value int) *Node {
	newNode := &Node{value: value, next: head}
	return newNode
}

func remove(head *Node, x int) *Node {
	current := head
	if current.value == x {
		head = current.next
		return head
	}

	previous := current
	for current != nil {
		if current.value == x {
			break
		}
		previous = current
		current = current.next
	}

	if current == nil {
		return head
	}
	previous.next = current.next

	return head
}


func main()  {
	head := &Node{
		value: 1,
		next:  nil,
	}
	addItem(head, &Node{value:2, next: nil})
	addItem(head, &Node{value:3})
	fmt.Println(getItem(head, 1))

	fmt.Println(search(head, 1))
	insertFront(head, 0)
	head = insertFront(head, 0)
	fmt.Println(getItem(head, 1))
	head = remove(head, 1)
	fmt.Println(head.next)
}
