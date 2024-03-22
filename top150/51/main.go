package main

import (
	"fmt"
	"sort"
)

func main() {
	v := [][]int{{1, 16}, {2, 8}, {1, 6}, {6, 12}, {4, 7}}
	//[[1,16],[2,8],[1,6],[6,12],[4,7]]
	answer := findMinArrowShots(v)
	fmt.Println(answer)
}
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	answer := 0
	fmt.Println(points)
	n := len(points)
	for i := 0; i < n; {
		min := points[i][1]
		for i < n-1 && min >= points[i+1][0] {
			if min > points[i+1][1] {
				min = points[i+1][1]
			}
			i++
		}
		i++
		answer++
	}
	return answer
}
