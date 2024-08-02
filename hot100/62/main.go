package main

import (
	"fmt"
)

func main() {
	res := solveNQueens(4)
	fmt.Println(res)
}
func solveNQueens(n int) [][]string {
	var chessboard = make([][]byte, n)
	var res [][]string
	for i := 0; i < n; i++ {
		chessboard[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chessboard[i][j] = '.'
		}
	}
	isValid := func(r, c int) bool {
		//竖着
		for i := 0; i < r; i++ {
			if chessboard[i][c] == 'Q' {
				return false
			}
		}

		// 左上对角线检查
		for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if chessboard[i][j] == 'Q' {
				return false
			}
		}
		// 右上对角线检查
		for i, j := r-1, c+1; i >= 0 && j <= n-1; i, j = i-1, j+1 {
			if chessboard[i][j] == 'Q' {
				return false
			}
		}
		//为何不检查左下和右下对角线
		//因为下方的行还未放置皇后，所以无需检查左下和右下方向。

		return true
	}
	var dfs func(r int)
	dfs = func(r int) {
		if r == n {
			var tmp = make([]string, n)
			for i := 0; i < n; i++ {
				tmp[i] = string(chessboard[i])
			}
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(r, i) {
				chessboard[r][i] = 'Q'
				dfs(r + 1)
				chessboard[r][i] = '.'
			}
		}
	}
	dfs(0)
	return res
}
