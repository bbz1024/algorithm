package main

func main() {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}
	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	rightSideView(&node1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var dfs func(node *TreeNode, level int)
	//level: 记录的是深度，先从右节点进入。若是深度与res的长度相同的话就是该层级的最右节点
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 0)
	return res
}
