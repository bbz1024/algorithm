package main

import (
	"fmt"
	"sort"
)

func main() {
	num := threeSum([]int{5, 6, 41, 2, 14, 0, 10, 7})
	fmt.Println(num)
}
func threeSum(nums []int) [][]int {
	answer := make([][]int, 0)
	sort.Ints(nums)
	for i, num := range nums {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if sum > 0 {
				l++
			} else if sum < 0 {
				r--
			} else {
				answer = append(answer, []int{num, nums[l], nums[r]})
				// 去重
				for r > l && nums[r] == nums[r-1] {
					r--
				}
				for r > l && nums[l] == nums[l+1] {
					l++
				}
				r--
				l++
			}
		}
	}
	return answer
}
