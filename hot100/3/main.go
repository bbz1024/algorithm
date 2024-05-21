package main

import (
	"fmt"
)

func main() {
	c := longestConsecutive([]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6})
	fmt.Println(c)
}
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
