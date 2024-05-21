package main

import "fmt"

func main() {
	res := isPalindrome(112)
	fmt.Println(res)
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	num := 0
	temp := x
	for temp != 0 {
		//相当于每遍历一次进位10
		num = num*10 + temp%10
		temp /= 10
	}
	return num == x
}
