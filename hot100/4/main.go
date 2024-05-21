package main

import "fmt"

func main() {
	arr := []int{0, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}
func moveZeroes(nums []int) {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != 0 {
			i++
			j++
			continue
		}
		for j < len(nums) && nums[j] == 0 {
			j++
		}
		if j >= len(nums) {
			return
		}
		nums[i] = nums[j]
		nums[j] = 0
		i++
		j++
	}
}
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
