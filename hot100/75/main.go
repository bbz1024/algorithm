package main

import (
	"fmt"
	"sort"
)

func main() {
	res := topKFrequent2([]int{-1, -1}, 1)
	fmt.Println(res)
}
func topKFrequent(nums []int, k int) []int {
	// 排序
	sort.Ints(nums)
	// 统计
	mp := make(map[int][]int)
	maxVal := 0
	for i := 0; i < len(nums); i++ {
		//	去重
		var cnt int
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			cnt++
			i++
		}
		if maxVal < cnt+1 {
			maxVal = cnt + 1
		}
		mp[cnt+1] = append(mp[cnt+1], nums[i])
	}
	var res = make([]int, 0, k)
	for i := maxVal; i > 0 && k > 0; {
		if v, ok := mp[i]; ok {
			res = append(res, v...)
			k -= len(v)
		}
		i--
	}
	return res
}

// -------------------- by heap --------------------

var mp = make(map[int]int)

// by heap
func topKFrequent2(nums []int, k int) []int {
	// cal num freq
	for _, num := range nums {
		mp[num]++
	}
	// 把每个num看做一个节点，mp作为存储节点的信息
	var setNum []int
	for num := range mp {
		setNum = append(setNum, num)
	}
}
