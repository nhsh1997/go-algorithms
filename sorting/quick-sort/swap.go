package quick_sort

func swap(list []int, i int, j int) {
	list[i], list[j] = list[j], list[i]
}
