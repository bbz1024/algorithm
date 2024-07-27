package main

import "fmt"

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
func letterCombinations(digits string) []string {
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var dfs func(startIdx int)
	var path string
	dfs = func(startIdx int) {
		if len(path) == len(digits) {
			res = append(res, path)
			return
		}
		str := mp[digits[startIdx]]
		for j := 0; j < len(str); j++ {
			path += string(str[j])
			dfs(startIdx + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
