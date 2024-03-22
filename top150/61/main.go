package main

import "fmt"

func main() {

	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Left = node2
	node1.Right = node3
	arr := make([]int, 10)
	fmt.Println(arr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, p.Left) && isSameTree(q.Left, q.Right)
}
