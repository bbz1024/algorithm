package main

import (
	"fmt"
	"sort"
)

func main() {
	//res := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	res := merge([][]int{{1, 4}, {4, 5}})
	fmt.Println(res)
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	if len(intervals) == 0 {
		return ans
	}
	ans = append(ans, intervals[0])
	if len(intervals) == 1 {
		return ans
	}
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		} else if ans[len(ans)-1][1] >= intervals[i][0] && ans[len(ans)-1][1] < intervals[i][1] {
			ans[len(ans)-1][1] = intervals[i][1]
		}
	}
	return ans
}
