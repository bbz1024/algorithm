package main

import "fmt"

func main() {
	//[2,3,4,5,6,7]
	res := isValid("([]{}")
	fmt.Println(res)
}
func isValid(s string) bool {
	var stack []byte
	mp := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mp[pop] != s[i] {
				return false
			}
		}

	}
	if len(stack) != 0 {
		return false
	}
	return true
}
