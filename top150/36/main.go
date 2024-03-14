package main

func main() {

}
func rotate(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
	}
	for i, rows := range matrix {
		for j, v := range rows {
			temp[j][n-1-i] = v
		}
	}
	copy(matrix, temp)
}
