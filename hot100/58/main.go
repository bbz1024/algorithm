package main

import (
	"fmt"
	"slices"
)

func main() {
	count := combinationSum([]int{2, 3, 5}, 8)
	fmt.Println(count)
	/*
		path	sum  startIdx
		2		2		0
		2,3		5		1
		2,3,6	11		2
		2,3		5		1



	*/
}
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var dfs func(startIdx int)
	var sum int

	dfs = func(startIdx int) {
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}
		if sum > target {
			return
		}
		for i := startIdx; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			dfs(i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}

	}
	dfs(0)
	return res
}
