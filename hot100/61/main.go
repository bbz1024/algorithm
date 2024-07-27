package main

import "fmt"

func main() {
	//分割回文串
	fmt.Println(partition("aab"))
}
func partition(s string) [][]string {
	var res [][]string
	var path []string
	isPalindrome := func(s string, i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	var dfs func(startIdx int)
	dfs = func(startIdx int) {
		if startIdx == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := startIdx; i < len(s); i++ {
			if isPalindrome(s, startIdx, i) {
				path = append(path, s[startIdx:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}
