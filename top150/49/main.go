package main

import (
	"fmt"
	"sort"
)

func main() {

	res := merge([][]int{{5, 5}, {1, 1}, {5, 7}, {5, 7}, {1, 1}, {3, 4}, {4, 4}, {0, 1}, {5, 5}, {1, 2}, {5, 5}, {0, 2}})
	fmt.Println(res)
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var answer [][]int
	for _, interval := range intervals {
		if answer == nil || answer[len(answer)-1][1] < interval[0] {
			answer = append(answer, interval)

		} else {
			if answer[len(answer)-1][1] < interval[1] {
				answer[len(answer)-1][1] = interval[1]
			}
		}
	}
	return answer
}
