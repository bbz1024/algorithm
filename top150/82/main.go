package main

// https://leetcode.cn/problems/number-of-islands/?envType=study-plan-v2&envId=top-interview-150
func main() {

}

func dfs(grid [][]byte, x, y int) {
	//越界了就说明不存在包围
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 'X' || grid[x][y] == '#' {
		//已经搜索过了，或者到达了边界
		return
	}
	grid[x][y] = '#'
	//上下左右
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || j == 0 || i == m-1 || j == n-1
			if isEdge && board[i][j] == 'O' {

				dfs(board, i, j)
			}
		}
	}
	//进行清理
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}
