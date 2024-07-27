package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	dup := subsetsWithDup([]int{1, 2, 2})
	fmt.Println(dup)
}

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(startidx int)

	backtrack = func(startidx int) {
		res = append(res, slices.Clone(path))

		for i := startidx; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
		return
	}
	backtrack(0)
	return res
}
func subsetsWithDup(nums []int) [][]int {
	var res [][]int

	var dfs func(startIdx int)
	var path []int
	sort.Ints(nums)
	dfs = func(startIdx int) {
		res = append(res, slices.Clone(path))
		for i := startIdx; i < len(nums); i++ {
			// 排序去重
			if i != startIdx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}

		return
	}
	dfs(0)
	return res
}
