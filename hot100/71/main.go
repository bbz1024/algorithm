package main

import (
	"fmt"
	"strings"
)

func main() {
	res := decodeString("3[a2[c]]")
	fmt.Println(res)
}
func decodeString(s string) string {
	var dfs func(s string, i int) (int, string)
	dfs = func(s string, i int) (int, string) {
		var res string
		var multi int
		for i < len(s) {
			if s[i] >= '0' && s[i] <= '9' {
				multi = multi*10 + int(s[i]-'0')
			} else if s[i] == '[' {
				resI, temp := dfs(s, i+1)
				i = resI // 跳过 ]
				res = res + strings.Repeat(temp, multi)
				multi = 0
			} else if s[i] == ']' {
				return i, res
			} else {
				res += string(s[i])
			}
			i++
		}
		return i, res
	}
	_, res := dfs(s, 0)
	return res

}
