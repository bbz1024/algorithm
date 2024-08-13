package main

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	//heights := []int{2, 4}
	res := largestRectangleArea(heights)
	println(res)
}

func largestRectangleArea(heights []int) int {
	var res int
	var stack []int
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		var left, mid, right int
		right = i
		for heights[i] < heights[stack[len(stack)-1]] {
			// 2
			mid = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left = stack[len(stack)-1]
			res = max(res, (right-left-1)*heights[mid])
		}
		stack = append(stack, i)
	}
	return res
}

/*
2, 1, 5, 6, 2, 3
stack []
i=0
-
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func largestRectangleArea2(heights []int) int {
	var res int
	heights = append(heights, 0)
	var newHeight = []int{0}
	newHeight = append(newHeight, heights...)
	for i := 1; i < len(newHeight)-1; i++ {
		// 左边第一个小于该元素的idx
		cur := newHeight[i]
		var leftMint, rightMin int
		for j := i - 1; j > 0; j-- {
			if cur > newHeight[j] {
				leftMint = j
				break
			}
		}
		for j := i + 1; j < len(newHeight); j++ {
			if cur > newHeight[j] {
				rightMin = j
				break
			}
		}
		weight := rightMin - leftMint - 1
		res = max(res, cur*weight)
	}
	return res
}
