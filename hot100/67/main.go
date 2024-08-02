package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(findMin(nums))
}
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	/*
		把 x 与最后一个数 nums[n−1] 比大小：

		如果 x>nums[n−1]，那么可以推出以下结论：
			nums 一定被分成左右两个递增段；
			第一段的所有元素均大于第二段的所有元素；
			x 在第一段。
			最小值在第二段。
			所以 x 一定在最小值的左边。
		如果 x≤nums[n−1]，那么 x 一定在第二段。（或者 nums 就是递增数组，此时只有一段。）
			x 要么是最小值，要么在最小值右边。

	*/
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]

}
