package main

import "fmt"

func main() {
	//10
	fmt.Println(0 ^ 0b1010)
	//fmt.Println(0b1010 & ^0b1010) // 0b0101 5
	//fmt.Println(0b1010 & 0b0101)
}

// 因为题中明确相同的数组只会出现1次，那么使用^异或，即相同为0,不同为1
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i] //
	}
	return ans
}
