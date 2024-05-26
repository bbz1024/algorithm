package main

import "fmt"

func main() {
	//ans := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	ans := maxArea([]int{1, 0, 0, 0, 0, 0, 0, 2, 2})
	fmt.Println(ans)
}
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := 0
	for i < j {
		if height[i] < height[j] {
			if ans < height[i]*(j-i) {
				ans = height[i] * (j - i)
			}
			i++
		} else {
			if ans < height[j]*(j-i) {
				ans = height[j] * (j - i)
			}
			j--
		}
	}
	return ans
}
