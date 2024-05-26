package main

import (
	"fmt"
)

func main() {
	res := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	//[3,3,2,5]

	fmt.Println(res)
}

func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	push := func(v int) {
		for len(queue) > 0 && nums[v] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, v)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[queue[0]]
	for i := k; i < n; i++ {
		push(i)
		//操过了范围
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		ans = append(ans, nums[queue[0]])

	}
	return ans
}

/*
// 超时
func maxSlidingWindow(nums []int, k int) []int {
	var getMax = func(nums []int) int {
		max := math.MinInt32
		for i := 0; i < len(nums); i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		return max
	}

	var ans []int
	l, r := 0, k

	for r <= len(nums) {
		ans = append(ans, getMax(nums[l:r]))
		l++
		r++
	}
	return ans
}

// 超时
func maxSlidingWindow2(nums []int, k int) []int {
	var getMax = func(nums []int) (int, int) {
		max := math.MinInt32
		idx := 0
		for i := 0; i < len(nums); i++ {
			if max < nums[i] {
				idx = i
				max = nums[i]
			}
		}
		return max, idx
	}
	var ans []int
	preMax := math.MinInt32
	preIdx := k
	l, r := 0, k
	for r <= len(nums) {
		//	preIdx 在k的有效范围内 r-preIdx
		if !(r-preIdx < k && preMax > nums[r-1]) {
			preMax, preIdx = getMax(nums[l:r])
		}
		ans = append(ans, preMax)
		l++
		r++
	}
	return ans
}

*/
