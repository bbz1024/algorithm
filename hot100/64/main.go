package main

import "fmt"

func main() {
	// write code here
	res := searchMatrix([][]int{{1, 3}}, 3)
	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	var res bool
	var search func(row int) bool
	search = func(row int) bool {
		rows := matrix[row]
		l, r := 0, len(rows)-1
		for l <= r {
			mid := (r-l)/2 + l
			if rows[mid] < target {
				l = mid + 1
			} else if rows[mid] > target {
				r = mid - 1
			} else {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if search(i) {
			res = true
			break
		}
	}
	return res
}
