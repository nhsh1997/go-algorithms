package main

import "fmt"

func main()  {
	fmt.Println(hIndex([]int{100, 5, 6, 7, 5}))
}

func hIndex(publications []int) int {
	citations := make(map[int]int)
	for i := 1; i <= len(publications); i++ {
		_, ok := citations[i]
		if !ok {
			citations[i] = 0
		}
	}

	hIndex := 0

	for k, _ := range citations {
		for i := 0; i < len(publications); i++ {
			if publications[i] >= k {
				citations[k]++
				if citations[k] >= k {
					hIndex = k
				}
			}
		}
	}
	fmt.Println(citations)

	return hIndex
}
