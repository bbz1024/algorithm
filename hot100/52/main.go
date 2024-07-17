package main

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
}

/*
模拟橘子腐烂过程的算法，其中 grid 是一个二维数组，表示一个网格，其中：
0 表示空格，
1 表示新鲜橘子，
2 表示腐烂橘子。
腐烂橘子会将周围四个方向的新鲜橘子在每一轮中都传染成腐烂状态。代码的目标是计算所有新鲜橘子完全腐烂所需的最小轮数，如果不可能则返回 -1。
*/
func orangesRotting(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
	var queue [][2]int
	count := 0
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			if grid[r][c] == 1 {
				count++
			} else if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}
	var round int
	for count > 0 && len(queue) > 0 {
		round++
		n := len(queue) // 层次遍历，遍历感染坏橘子的周边
		for i := 0; i < n; i++ {
			r, c := queue[0][0], queue[0][1]
			queue = queue[1:] // 移除已处理的元素
			if r-1 >= 0 && grid[r-1][c] == 1 {
				grid[r-1][c] = 2
				count--
				queue = append(queue, [2]int{r - 1, c})
			}
			if r+1 < M && grid[r+1][c] == 1 {
				grid[r+1][c] = 2
				count--
				queue = append(queue, [2]int{r + 1, c})
			}
			if c-1 >= 0 && grid[r][c-1] == 1 {
				grid[r][c-1] = 2
				count--
				queue = append(queue, [2]int{r, c - 1})
			}
			if c+1 < N && grid[r][c+1] == 1 {
				grid[r][c+1] = 2
				count--
				queue = append(queue, [2]int{r, c + 1})
			}
		}
	}
	if count > 0 {
		return -1
	}
	return round
}
