package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func total(num int, node *TreeNode) int {
	if node == nil {
		return 0
	}
	num = num*10 + node.Val
	if node.Left == nil && node.Right == nil {

		return num
	}
	return total(num, node.Right) + total(num, node.Left)
}
func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return total(0, root)
}
