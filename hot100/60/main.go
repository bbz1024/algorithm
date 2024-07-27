package main

import "fmt"

func main() {
	var board = [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	res := exist(board, "cdba")
	fmt.Println(res)
}
func exist(board [][]byte, word string) bool {
	var dfs func(i, j, idx int)
	var visited = make([][]bool, len(board))
	var res bool
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	dfs = func(i, j, idx int) {
		//		check bound
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return
		}
		if visited[i][j] {
			return
		}

		if board[i][j] != word[idx] {
			return
		}
		// 只有相同才进行标记，也就是找到了才进行标记，否则会出现未匹配也进行标记，
		visited[i][j] = true
		// 收结果了
		if idx == len(word)-1 {
			res = true
			return
		}
		dfs(i+1, j, idx+1)
		dfs(i-1, j, idx+1)
		dfs(i, j+1, idx+1)
		dfs(i, j-1, idx+1)
		visited[i][j] = false // 回溯
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				dfs(i, j, 0)
			}
		}
	}
	return res
}
