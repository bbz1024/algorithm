package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})
	fmt.Println(res)
}
func summaryRanges(nums []int) []string {
	var answer []string
	var str strings.Builder
	n := len(nums)
	for i := 0; i < n; {
		str.WriteString(strconv.Itoa(nums[i]))
		j := i
		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		if j != i {
			str.WriteString("->")
			str.WriteString(strconv.Itoa(nums[i]))
		}
		answer = append(answer, str.String())
		str.Reset()
		i++
	}
	return answer
}
