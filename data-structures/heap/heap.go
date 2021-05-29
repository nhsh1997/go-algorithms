package main

import "fmt"

type IHeap interface {
	Insert(x int)
	DeleteMax() int
	LeftChild(i int) int
	RightChild(i int) int
	Parent(i int) int
	fixUp(i int)
	fixDown()
	swap(x, y int)
}

type Heap struct {
	items []int
	n int
}

func (h *Heap) fixDown() {
	i := 0
	for i <= h.n - 1 {
		j := h.LeftChild(i)
		if j <= h.n - 1 {
			if h.items[j] < h.items[j + 1]{
				j++
			}
			h.swap(i, j)
		}
		i = j
	}
}

func (h *Heap) DeleteMax() int {
	if h.n == 0 {
		return -1
	}
	max := h.items[0]
	h.swap(0, h.n-1)
	h.items[h.n-1] = 0
	h.n --
	h.fixDown()
	return max
}

func (h *Heap) swap(x, y int) {
	temp := h.items[x]
	h.items[x] = h.items[y]
	h.items[y] = temp
}

func (h *Heap) fixUp(i int) {
	for h.Parent(i) >= 0 && h.items[h.Parent(i)] < h.items[i] {
		h.swap(i, h.Parent(i))
		i = h.Parent(i)
	}
}

func (h *Heap) Insert(x int) {
	h.n++
	if len(h.items) < h.n {
		h.items = append(h.items, x)
	} else {
		h.items[h.n -1] = x
	}
	h.fixUp(h.n-1)
}

func (h *Heap) LeftChild(i int) int {
	return 2*i + 1
}

func (h *Heap) RightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) Parent(i int) int {
	return (i - 1)/2
}

func NewHeap() IHeap {
	return &Heap{
		items: []int{},
	}
}

func main() {
	heap := NewHeap()
	heap.Insert(5)
	//fmt.Println(heap)
	heap.Insert(19)
	//fmt.Println(heap)
	heap.Insert(18)
	//fmt.Println(heap)
	heap.Insert(30)
	//fmt.Println(heap)
	heap.Insert(21)
	//fmt.Println(heap)
	heap.Insert(25)
	//fmt.Println(heap)
	heap.Insert(2)
	//fmt.Println(heap)
	heap.Insert(4)
	//fmt.Println(heap)
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap)
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap)
	heap.Insert(100)
	fmt.Println(heap)
	heap.Insert(80)
	fmt.Println(heap)
	heap.Insert(99)
	fmt.Println(heap)
}
