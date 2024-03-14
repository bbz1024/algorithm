package main

/*
https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-interview-150
旋转矩阵
*/
func main() {

}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var (
		rows, columns            = len(matrix), len(matrix[0])
		answer                   = make([]int, rows*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			answer[index] = matrix[top][i]
			index++
		}
		for i := top + 1; i <= bottom; i++ {
			answer[index] = matrix[i][right]
			index++
		}
		if left < right && top < bottom {
			for i := right - 1; i > left; i-- {
				answer[index] = matrix[bottom][i]
				index++
			}
			for i := bottom; i > top; i-- {
				answer[index] = matrix[i][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--

	}
	return answer
}
