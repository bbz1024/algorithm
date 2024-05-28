package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 6)
	fmt.Println(arr)

}
func rotate(nums []int, k int) {
	k %= len(nums)
	var temp int
	for i := 0; i < k; i++ {
		temp = nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}
func rotate2(nums []int, k int) {
	//slices.Reverse(nums)
	//slices.Reverse(nums[k%len(nums):])
	//slices.Reverse(nums[0:k%len(nums)])

}
