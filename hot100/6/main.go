package main

import "sort"

func main() {

}
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				//	等于0，收货结果
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 针对l和r进行去重
				//l去重
				for r > l && nums[l] == nums[l+1] {
					l++
				}
				//r 去重
				for r > l && nums[r] == nums[r-1] {
					r--

				}
				l++
				r--
			}
		}
	}
	return res
}
