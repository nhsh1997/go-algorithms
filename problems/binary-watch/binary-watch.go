package main

import "fmt"


func main()  {
	fmt.Println(readBinaryWatch(8))
}

type Watch struct {
	Hours [][]int
	Minutes [][]int
}

func (w *Watch) convertToString() []string {
	var listString []string
	for _, hour := range w.Hours {
		h := sum(hour)
		for _, minute := range w.Minutes {
			var minuteString string
			m := sum(minute)
			if h < 12 && m < 60 {
				if m < 10 {
					minuteString = fmt.Sprintf("0%d", m)
				} else {
					minuteString = fmt.Sprintf("%d", m)
				}
				listString = append(listString, fmt.Sprintf("%d:%s", h, minuteString))
			}
		}
	}
	return listString
}

func sum(list []int) int {
	total := 0

	for _, value := range list {
		total += value
	}

	return total
}


func readBinaryWatch(turnedOn int) []string {
	if turnedOn < 0 || turnedOn > 10 {
		return []string{}
	}

	var result []string

	hours := []int{8, 4, 2, 1}
	minutes := []int{32, 16, 8, 4, 2, 1}

	n := turnedOn

	if turnedOn > len(hours) {
		n = len(hours)
	}

	for i := 0; i <= n ; i++  {
		j := turnedOn - i
		hoursSubsets := new([][]int)
		findSubset(hours, []int{}, 0, i, hoursSubsets)
		minutesSubsets := new([][]int)
		findSubset(minutes, []int{}, 0, j, minutesSubsets)
		watch := Watch{
			Hours:   *hoursSubsets,
			Minutes: *minutesSubsets,
		}
		wString := watch.convertToString()
		result = append(result, wString...)
	}

	return result
}

// findSubset using backtrack algorithm to find the subsets of an array number.
// @param {[]int} nums is array of natural numbers
// @param {[]int} list is the recursive subset
// @param {int} target is number of elements in the subset
// @param {[][]int{}} array of subsets
func findSubset(nums, list []int, start, target int, result *[][]int) {

	if len(list) == target {
		*result = append(*result, list)
	}

	for i := start; i < len(nums); i++  {
		subset := copy(list)
		subset = append(subset, nums[i])
		findSubset(nums, subset, i + 1, target, result)
	}
}

func copy(src []int) []int {
	var result []int
	return append(result, src...)
}



