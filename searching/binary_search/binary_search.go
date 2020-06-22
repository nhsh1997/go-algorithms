package binary_search

func BinarySearch(list []int, left int, right int, x int) int {
	if right >= left {
		mid := left + (right - left) / 2

		if list[mid] == x {
			return mid
		} else {
			if mid == 0 {
				return -1
			}
		}

		if list[mid] > x {
			return BinarySearch(list, left, mid - left, x)
		}

		return  BinarySearch(list, mid + 1, right, x)
	}
	return -1
}


/*func main()  {
	list := []int{2, 4, 6, 9, 10}
	N := len(list)
	x := 4
	result := BinarySearch(list, 0, N - 1, x)

	if result != -1 {
		fmt.Println(fmt.Sprintf("Index of ", x, " is ", result))
	} else {
		fmt.Println(fmt.Sprintf("Not found ", x))
	}
}
*/