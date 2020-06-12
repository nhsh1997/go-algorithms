package quick_sort

func QuickSort(list []int, ascending bool) {
	sort(list, 0, len(list)-1, ascending)
}
