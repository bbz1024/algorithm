package main

import "sort"

func main() {

}
func insert(intervals [][]int, newInterval []int) (answer [][]int) {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 可以进行先插入后排序在调整区间
	for _, interval := range intervals {
		if len(answer) == 0 || answer[len(answer)-1][1] < interval[0] {
			answer = append(answer, interval)
		} else {
			if answer[len(answer)-1][1] < interval[1] {
				answer[len(answer)-1][1] = interval[1]
			}
		}
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	l := len(intervals)
	i := 0 // 记录插入的位置
	//筛选出小于新的[0]的左部分，也可以是左部分不存在重合
	for i < l && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	//更新存在重合的部分
	for i < l && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	//插入右边不存在重叠的部分
	for i < l {
		res = append(res, intervals[i])
		i++
	}
	return res
}
