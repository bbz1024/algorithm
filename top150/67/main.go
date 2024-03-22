package main

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Left = &TreeNode{Val: 6}
	root.Right.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right.Right = &TreeNode{Val: 8}
	PreOrder(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodes []*TreeNode

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	nodes = append(nodes, root)
	PreOrder(root.Left)
	PreOrder(root.Right)
	root.Left = nil

}

func flatten(root *TreeNode) {
	PreOrder(root)
	for i := 1; i < len(nodes); i++ {

		nodes[i-1].Right = nodes[i]
	}
}
