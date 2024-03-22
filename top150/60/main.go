package main

import "fmt"

func main() {

	tree := TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 3}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Right.Left = &TreeNode{Val: 5}
	tree.Right.Right = &TreeNode{Val: 6}
	fmt.Println(maxDepth(&tree))

	/*
					1
			2				3
		4				5		6

	*/
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0 // 从最后一个加起
	} else {
		l := maxDepth(root.Left)
		r := maxDepth(root.Right)
		if l > r {
			return l + 1
		}
		return r + 1
	}
}
