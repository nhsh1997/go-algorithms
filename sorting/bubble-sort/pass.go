package bubble_sort

func pass(list []int, ascending bool){
	N := len(list)
	for i := 0; i < N ; i ++ {
		firstElement := list[i]
		secondElement := list[i + 1]
		if ascending == true {
			 if firstElement < secondElement {
			 	swap(list, i)
			 }
		} else {
			if firstElement > secondElement {
				swap(list, i)
			}
		}
	}
}