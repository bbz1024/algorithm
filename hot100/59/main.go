package main

import "fmt"

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
func generateParenthesis(n int) []string {

	var res []string
	var path string
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if right > left {
			return
		}
		if left == right && left == n {
			res = append(res, path)
			return
		}
		if left < n {
			path = path + "("
			dfs(left+1, right)
			path = path[:len(path)-1]
		}
		if right < n {
			path = path + ")"
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return res
}
