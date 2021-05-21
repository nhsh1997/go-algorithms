package main

import "fmt"

type Watch struct {
	Hours map[int]int
	Minutes map[int]int
}

func (w *Watch) toString() string {
	totalHours := 0
	totalMinutes := 0
	for _, v := range w.Hours {
		totalHours += v
	}

	for _, v := range w.Minutes {
		totalMinutes += v
	}

	hourStr := fmt.Sprintf("%d", totalHours)
	if totalHours < 10 {
		hourStr = fmt.Sprintf("0%d", totalHours)
	}

	minuteStr := fmt.Sprintf("%d", totalMinutes)
	if totalMinutes < 10 {
		minuteStr = fmt.Sprintf("0%d", totalMinutes)
	}

	return fmt.Sprintf("%s:%s", hourStr, minuteStr)


}

func main()  {
	readBinaryWatch(2)
}

func readBinaryWatch(turnedOn int) []string {
	hours := []int{0, 8, 4, 2, 1}
	minutes := []int{32, 16, 8, 4, 2, 1, 0}
	result := make(map[int][]Watch)
	result[0] = []Watch{}
	backtrack(turnedOn, hours, minutes, nil, result)
	fmt.Println(result[0])

	var resultStr []string
	for _, watch := range result[0] {
		resultStr = append(resultStr, watch.toString())
	}
	fmt.Println(resultStr)
	return nil
}

func backtrack(turnedOn int, hours []int, minutes[]int, watch *Watch, mapResult map[int][]Watch) {
	//Check if turned on numbers of led yet
	if turnedOn == 0 {
		return
	} else {
		// loop hours
		for i, hour := range hours {
			// init new watch
			newWatch := Watch{
				Hours: map[int]int{},
				Minutes: map[int]int{},
			}

			// check this is recursive
			if watch != nil {
				for hour := range watch.Hours  {
					newWatch.Hours[hour] = hour
				}
				for minute := range watch.Minutes  {
					newWatch.Minutes[minute] = minute
				}
			}

			// check is hours[i] is exist in newWatch.Hours?
			_, ok := newWatch.Hours[hour]
			if !ok {
				newWatch.Hours[hour] = hour
				mapResult[0] = append(mapResult[0], newWatch)
				// backtracking without hours[i]
				backtrack(turnedOn - 1, hours[(i + 1):], minutes, &newWatch, mapResult)
			}
		}
	}

}