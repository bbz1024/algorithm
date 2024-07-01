package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var cnt int
	var res int
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		preOrder(root.Left)
		cnt++
		if cnt == k {
			res = root.Val
		}
		preOrder(root.Right)
	}
	preOrder(root)
	return res
}
