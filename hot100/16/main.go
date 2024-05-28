package main

func main() {
	productExceptSelf([]int{1, 2, 3, 4})
}

// 前缀和
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	end := make([]int, len(nums))

	pre[0] = 1
	end[len(end)-1] = 1

	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		end[i] = end[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = pre[i] * end[i]
	}
	return nums
}
