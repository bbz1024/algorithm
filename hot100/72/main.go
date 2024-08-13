package main

import "fmt"

func main() {
	//res := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	//res := dailyTemperatures([]int{40, 10, 10, 30})
	//res := dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70})
	//res := dailyTemperatures([]int{3, 4, 5, 6})
	res := dailyTemperatures([]int{77, 41, 77, 41, 41, 77})
	fmt.Println(res)
}
func dailyTemperatures(temperatures []int) []int {
	var stack []int // temperatures, index
	var res = make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if i < len(temperatures)-1 && temperatures[i] < temperatures[i+1] {
			res[i] = 1
			continue
		}
		stack = append(stack, i)
	}
	return res
}
