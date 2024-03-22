package main

import (
	"strings"
)

func main() {
	//simplifyPath("/home/")
	//simplifyPath("/home//foo/")
	simplifyPath("/../")
	//simplifyPath("/a/./b/../../c/")
}
func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	var stack []string
	for _, dir := range dirs {
		if dir == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else if dir == "." || dir == "" {
			continue
		} else {
			stack = append(stack, dir)
		}
	}
	return "/" + strings.Join(stack, "/")
}
