package main

import (
	"fmt"
	"math"
)

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(tree)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var res = math.MinInt32
	var sumPath func(root *TreeNode) int
	sumPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := sumPath(root.Left)
		right := sumPath(root.Right)
		sum := root.Val + max(left, 0) + max(right, 0)
		ret := root.Val + max(0, max(right, left))
		res = max(res, max(sum, ret))
		return ret
	}
	sumPath(root)
	return res

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}
