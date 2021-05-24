package main

import "fmt"

type IStack interface {
	push(int)
	pop() int
	top() int
	isEmpty() bool
}

type Stack struct {
	items []int
	topIndex int
	maxSize int
}

func (s *Stack) push(x int) {
	if s.topIndex == s.maxSize {
		return
	} else {
		s.items[s.topIndex] = x
		s.topIndex++
	}
}

func (s *Stack) pop() int {
	value := s.items[s.topIndex - 1]
	s.items[s.topIndex - 1]  = 0
	s.topIndex--
	return value
}

func (s *Stack) top() int {
	return s.items[s.topIndex - 1]
}

func (s *Stack) isEmpty() bool {
	return s.topIndex == 0
}

func newStack(maxSize int) IStack {
	items := make([]int, maxSize)
	return &Stack{
		items: items,
		topIndex: 0,
		maxSize: maxSize,
	}
}

func main()  {
	stk := newDynamicStack()
	stk.push(1)
	stk.push(2)
	fmt.Println(stk.pop())
	fmt.Println(stk)
	fmt.Println(stk.pop())
	fmt.Println(stk)
	stk.push(2)
}
