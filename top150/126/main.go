package main

import "fmt"

func main() {
	//1011,0001
	res := hammingWeight(2147483647)
	fmt.Println(res)
}
func hammingWeight(n int) int {
	c := 0
	for ; n > 0; n = n >> 1 {
		if n&1 == 1 {
			c++
		}

	}
	return c
}
