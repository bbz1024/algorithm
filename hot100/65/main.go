package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	fmt.Println(searchRange(nums, 3))
}
func searchRange(nums []int, target int) []int {
	var res = make([]int, 2)
	res[0] = -1
	res[1] = -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			// 是否存在前后
			// 最坏情况下O（n）
			res[0], res[1] = mid, mid
			m := mid
			// 前去重
			for m > 0 && nums[m-1] == target {
				res[0] = m - 1
				m--
			}
			// 后去重
			for m < len(nums)-1 && nums[m+1] == target {
				res[1] = m + 1
				m++
			}

			break
		}
	}
	return res
}

// 闭区间写法
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 闭区间 [left, right]
	for left <= right {           // 区间不为空
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1 // 范围缩小到 [mid+1, right]
		} else {
			right = mid - 1 // 范围缩小到 [left, mid-1]
		}
	}
	return left
}

func searchRange2(nums []int, target int) []int {
	start := lowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 如果 bound 存在，那么 end 必定存在 ，寻找一个刚好大于target 的数，那么他的前一个必定是target的下标
	end := lowerBound(nums, target+1) - 1
	return []int{start, end}
}
