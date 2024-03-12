package main

import (
	"math"
)

func main() {
	candy([]int{1, 2, 87, 87, 87, 2, 1})
	// left ：[1 2 3 1 1 1 1]
	//			[1 2 3 1 3 2 1]

}

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < len(left); i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(left) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] {
			right[i] = right[i+1] + 1
		}
	}
	count := 0
	for i := 0; i < len(ratings); i++ {
		count += int(math.Max(float64(left[i]), float64(right[i])))
	}
	return count

}
