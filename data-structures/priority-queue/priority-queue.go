package priority_queue

type IPriorityQueue interface {
	Insert(value int)
	DeleteMax() int
	DeleteMin() int
}
