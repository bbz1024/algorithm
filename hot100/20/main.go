package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(arr)
	fmt.Println(arr)
}
func rotate(matrix [][]int) {
	// x,y => y,n-x
	n := len(matrix)
	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, n)
		copy(tmp[i], matrix[i])
	}

	// 开始转换
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-1-i] = tmp[i][j]
		}
	}
}
