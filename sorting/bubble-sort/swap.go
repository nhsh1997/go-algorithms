package bubble_sort

func swap(list []int, firstIndex int)  {
	list[firstIndex], list[firstIndex + 1] = list[firstIndex + 1], list[firstIndex]
}