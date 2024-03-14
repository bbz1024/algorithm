package main

func main() {
	isValidSudoku(nil)

}

func isValidSudoku(board [][]byte) bool {
	var box [3][3][9]uint8
	var row [9][9]uint8
	var column [9][9]uint8
	for x, inner := range board {
		for y, v := range inner {
			if v == '.' {
				continue
			}
			idx := v - '0'
			row[x][idx]++    // 横
			column[y][idx]++ // 竖
			box[x/3][y/3][idx]++
			if row[x][idx] > 1 || column[y][idx] > 1 || box[x/3][y/3][idx] > 1 {
				return false
			}
			//	if c!=nil {
		}
	}
	return true
}
