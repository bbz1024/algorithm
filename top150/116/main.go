package main

import "fmt"

func main() {
	fmt.Println(searchMatrix2([][]int{{2}}, 0))
}
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

/*
在二分查找算法中计算中间索引时，需要加上 left 的原因是为了确保得到的是左闭区间 [left, right] 内的中间位置。
由于我们通常从数组的第一个元素开始搜索，并逐步缩小搜索范围，因此需要根据当前的 left 和 right 边界来确定下一个要检查的中间元素。
在表达式 mid := left + (right - left) / 2 中：

	right - left 计算了区间长度。
	将区间长度除以 2 后，得到的是区间长度的一半（向下取整）。
	加上 left 后，就得到了区间 [left, right] 的中间索引位置。
	如果不加 left，得到的就是左开区间 (left, right) 的中间位置，这在数组索引的情况下会导致访问越界或者错过可能相等的目标值
*/
func search(arr []int, target int) bool {

	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	x := 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > target {
			break
		}
		x = i
	}
	return search(matrix[x], target)
}
