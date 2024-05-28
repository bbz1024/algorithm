package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxSubArray3([]int{-1, 1, 2, 1})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	var ans = math.MinInt32
	left, right := 0, 0
	for right < len(nums) {
		sum := 0
		for i := 0; i < right-left+1; i++ {
			sum += nums[left : right+1][i]
		}
		if nums[right] > ans && nums[right] > sum {
			ans = nums[right]
			left = right
		} else if nums[right] > sum {
			left = right
		}

		if sum > ans {
			ans = sum
		}
		right++
	}
	return ans
}

func maxSubArray2(nums []int) int {
	maxCurrent, maxGlobal := nums[0], nums[0]
	for _, num := range nums[1:] {
		maxCurrent = max(num, maxCurrent+num)
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}
	return maxGlobal
}

// Kadane 算法  https://blog.csdn.net/weixin_43764974/article/details/134513506
/*
Kadane's Algorithm 是解决“最大子序列和”（Maximum Subarray Problem）问题的一种高效算法，
它可以在单遍历数组的过程中找到数组中具有最大和的连续子数组。该算法的时间复杂度为 O(n)，其中 n 是数组的长度。以下是 Kadane 算法的基本步骤：
	1. 初始化两个变量，max_current 用于记录当前遍历到的子数组的最大和，初始值为数组的第一个元素；max_global 用于记录全局的最大和，也初始化为第一个元素。
	2. 遍历数组的其余部分（从第二个元素开始）：
		对于每个元素 num：
			计算 max_current 是否可以通过加上 num 增大。如果可以，就更新 max_current 为 max_current + num；否则，将 max_current 重置为 num，因为当前子数组的和已经不可能增加。
			比较 max_current 和 max_global，如果 max_current 更大，则更新 max_global 为 max_current。
	3. 在遍历结束后，max_global 将包含整个数组中的最大子序列和。
*/
func maxSubArray3(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		// 加上一个还大于num[i] 说明是渐增的
		if nums[i-1]+nums[i] > nums[i] {
			//前缀和
			nums[i] += nums[i-1]
		}

		ans = max(ans, nums[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
