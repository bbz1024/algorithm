package main

func main() {
	//setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
	setZeroes2([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})
}
func setZeroes(matrix [][]int) {

	x := len(matrix)
	y := len(matrix[0])
	mp := make([][]bool, x)

	for i := 0; i < x; i++ {
		mp[i] = make([]bool, y)
		for j := 0; j < y; j++ {
			if matrix[i][j] == 0 {
				mp[i][j] = true
			}
		}
	}
	toZero := func(x1, y1 int) {
		// 横
		for i := 0; i < y; i++ {
			matrix[x1][i] = 0
		}
		// 竖
		for i := 0; i < x; i++ {
			matrix[i][y1] = 0
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if mp[i][j] {
				toZero(i, j)
			}
		}
	}
}
func setZeroes2(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, cols := range matrix {
		for j, v := range cols {
			if v == 0 {
				// 进行标记这行需要全部为0
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, cols := range matrix {
		for j, _ := range cols {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
