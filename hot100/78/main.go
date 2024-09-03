package main

import "fmt"

func main() {
	var arr []int
	arr = []int{2, 3, 1, 1, 4}
	//arr = []int{3, 2, 1, 0, 4}
	jump := canJump(arr)
	fmt.Println(jump)

}
func canJump(nums []int) bool {
	mx := 0 // 最大覆盖范围
	for i, jump := range nums {
		if i > mx { // 无法到达 i
			return false
		}
		mx = max(mx, i+jump) // 从 i 最右可以跳到 i + jump
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
