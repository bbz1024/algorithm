package main

import "fmt"

func main() {
	ans := lengthOfLongestSubstring("pwwkew")
	fmt.Println(ans)
}
func lengthOfLongestSubstring(s string) int {
	var hp = make(map[byte]struct{})
	maxSubStringLen := 0
	start, end := 0, 0
	for end < len(s) {
		if _, ok := hp[s[end]]; ok {
			delete(hp, s[start])
			start++
		} else {
			hp[s[end]] = struct{}{}
			end++
		}
		if maxSubStringLen < end-start {
			maxSubStringLen = end - start
		}
	}
	return maxSubStringLen
}
