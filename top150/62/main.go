package main

import "fmt"

func main() {

	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node1.Left = node2
	node1.Right = node3
	res := invertTree(node1)
	fmt.Println(res)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从上到下进行递归交换两个子节点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := &TreeNode{Left: root.Left}
	cur := temp.Left
	temp2 := &TreeNode{Right: root.Right}
	cur2 := temp2.Right
	root.Right = invertTree(cur)
	root.Left = invertTree(cur2)
	return root
}
