package quick_sort

func sort(list []int, start int, end int, ascending bool){
	if (end - start) < 1 {
		return
	}

	pivot := list[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if ascending == true {
			if list[i] < pivot {
				swap(list, splitIndex, i)
				splitIndex++
			}
		} else {
			if list[i] > pivot {
				swap(list, splitIndex, i)
				splitIndex++
			}
		}
	}

	list[end] = list[splitIndex]
	list[splitIndex] = pivot

	sort(list, start, splitIndex-1, ascending)
	sort(list, splitIndex+1, end, ascending)
}
