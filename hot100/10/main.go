package main

import "fmt"

func main() {
	res := subarraySum3([]int{0, 0, 0, 0, 1}, 1)
	fmt.Println(res)
}

// 这里只能解决正整数的
func subarraySum(nums []int, k int) int {
	l, r := 0, 0
	ans := 0
	for r < len(nums) {
		sum := 0
		if l > r {
			return ans
		}
		for i := l; i <= r; i++ {
			sum += nums[i]
		}
		if sum < k {
			r++
		} else if sum > k {
			l++
		} else {
			ans++
			r++
			l++
		}
	}
	return ans
}
func subarraySum2(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和+哈希表
func subarraySum3(nums []int, k int) int {
	count := 0
	mp := make(map[int]int) // 前置和 / 出现的次数
	pre := 0
	mp[0] = 1 // 处理 刚好nums[i]==k，从而跳过了0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if v, ok := mp[pre-k]; ok {
			count += v // 前面pre-k的值出现的次数
		}
		//处理多个连续的数
		mp[pre] += 1
	}
	return count
}
