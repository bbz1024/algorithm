package main

import "fmt"

func main() {
	/*
		[[1,2,3,4,5]
		[6,7,8,9,10]
		[11,12,13,14,15]
		[16,17,18,19,20]
		[21,22,23,24,25]
		]
	*/
	res := searchMatrix([][]int{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15}, []int{16, 17, 18, 19, 20}, []int{21, 22, 23, 24, 25}}, 19)
	fmt.Println(res)
}
func binarySearch(arr []int, target int) (idx int, exist bool) {
	left, right := 0, len(arr)-1
	for left <= right {
		//mid := left + (right-left)/2 // 在计算h时避免溢出
		mid := int(uint(left+right) >> 1) // 在计算h时避免溢出
		//h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if arr[mid] == target {
			return mid, true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left, false
}

func searchMatrix(matrix [][]int, target int) bool {
	for _, arr := range matrix {
		_, exist := binarySearch(arr, target)
		if exist {
			return true
		}
	}
	return false
}
