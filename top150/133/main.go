package main

import "fmt"

func main() {
	fmt.Println(mySqrt(8))
}

// 二分法
func mySqrt(x int) int {
	l, r := 0, x
	answer := -1
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			answer = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return answer
}
