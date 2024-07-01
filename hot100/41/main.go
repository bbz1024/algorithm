package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var depth func(root *TreeNode) int
	if root == nil {
		return res
	}
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 左孩子的深度
		l := depth(root.Left)
		// 右孩子的深度
		r := depth(root.Right)
		// 更新ans +1 是因为加上负父节点
		res = int(math.Max(float64(res), float64(r+l+1)))
		return int(math.Max(float64(l), float64(r))) + 1
	}
	depth(root)
	return res - 1
}
