package main

import "fmt"

func main() {
	insert := searchInsert([]int{1, 2, 3, 4, 5, 6}, 8)
	fmt.Println(insert)
}

// 1,2,3,4,5,6

// 0,5  tag=4 mid=2
// l=mid+1=3,5

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := len(nums)
	for l <= r {
		mid := (r + l) / 2
		if target <= nums[mid] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
