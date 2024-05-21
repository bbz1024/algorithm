package main

// https://leetcode.cn/problems/number-of-islands/?envType=study-plan-v2&envId=top-interview-150
func main() {

}
func dfs(grid [][]byte, x, y int) {
	if bound(grid, x, y) {
		return
	}
	//是否是d岛屿
	if grid[x][y] != '1' {
		return
	}
	//标记为已经走过的，
	grid[x][y] = '2'
	//上下左右
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}

// 是否越界
func bound(grid [][]byte, x, y int) bool {
	row := len(grid) - 1
	col := len(grid[0]) - 1
	return x > row || x < 0 || y > col || y < 0
}

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	if len(grid) == 1 && len(grid[0]) == 0 {
		return 0
	}
	answer := 0
	for x, g := range grid {
		for y, b := range g {
			//只要存在’1‘,就说明还存在岛屿，不断更新岛屿范围
			if b == '1' {
				answer++
				dfs(grid, x, y)
			}
		}
	}
	return answer
}
