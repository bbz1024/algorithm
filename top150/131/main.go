package main

import "fmt"

func main() {
	res := plusOne2([]int{9})
	fmt.Println(res)
}
func plusOne(digits []int) []int {
	n := len(digits)
	var answer = make([]int, 0, n)
	add := true
	for i := n - 1; i >= 0; i-- {
		if add {
			digits[i] += 1
			add = false
		}
		if digits[i] == 10 {
			digits[i] = 0
			if i != 0 {
				digits[i-1] += 1
			}
		}
		answer = append(answer, digits[i])
	}
	if answer[n-1] == 10 || answer[n-1] == 0 {
		answer[n-1] = 0
		answer = append(answer, 1)
	}
	t := len(answer)
	for i := 0; i < t/2; i++ {
		answer[i], answer[t-1-i] = answer[t-1-i], answer[i]
	}
	return answer
}
func plusOne2(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
