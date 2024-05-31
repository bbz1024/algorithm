package main

import "fmt"

func main() {
	res := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(res)
}
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		rows, columns            = len(matrix), len(matrix[0])
		ans                      = make([]int, 0, rows*columns)
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	for left <= right && top <= bottom {
		//	左到右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}

		//	从上到下 top+1 避免取到重复的值
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}

		if left < right && top < bottom {
			//右到左
			for i := right - 1; i > left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			// 下到上
			for i := bottom; i > top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}
