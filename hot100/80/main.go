package main

import "fmt"

func main() {
	//res := partitionLabels("ababcbacadefegdehijhklij")
	res := partitionLabels("eaaaabaaec")
	fmt.Println(res)
}
func partitionLabels(s string) []int {
	var res []int
	var maxAppear int
	var start int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[j] == s[i] {
				maxAppear = max(maxAppear, j)
			}
		}
		if i == maxAppear {
			res = append(res, i+1-start)
			start = i + 1
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
