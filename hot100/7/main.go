package main

func main() {

}
func trap(height []int) int {
	var stack []int
	ans := 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(h, height[left]) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
