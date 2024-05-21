package main

func main() {
	//c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//res := searchInsert(c, 9)
	//fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
