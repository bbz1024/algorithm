package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(30))
}

// 会出现溢出
func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	num := jj(n)
	answer := 0
	for num%10 == 0 {
		num /= 10
		answer++
	}
	return answer
}

func jj(num int) int {
	if num == 1 {
		return 1
	}
	return num * jj(num-1)
}
