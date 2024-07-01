package main

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	println(numIslands(grid))
}

func dfs(grid [][]byte, x, y int) {
	if bound(grid, x, y) {
		return
	}
	//是否是d岛屿
	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	//
	dfs(grid, x-1, y)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x, y+1)
}
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
			if b == '1' {
				answer++
				dfs(grid, x, y)
			}
		}
	}
	return answer
}
