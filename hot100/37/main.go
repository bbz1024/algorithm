package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	//	 queue
	var res []int
	if root == nil {
		return res
	}
	var inOrderL func(node *TreeNode)
	inOrderL = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrderL(node.Left)
		res = append(res, node.Val)
		inOrderL(node.Right)
	}
	inOrderL(root)
	return res
}
