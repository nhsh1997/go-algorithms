package quick_sort

func QuickSort(list []int, ascending bool) {
	sort(list, 0, len(list)-1, ascending)
}


func quickSortRange(arr []int, first, last int) {
	if last <= first {
		return
	}
	pivot := arr[first]
	pos := last
	for i := last; i > first; i-- {
		if arr[i] > pivot {
			swap(arr, pos, i)
			pos--
		}
		swap(arr, first, pos)
	}
	quickSortRange(arr, first, pos - 1)
	quickSortRange(arr, pos + 1, last)
}