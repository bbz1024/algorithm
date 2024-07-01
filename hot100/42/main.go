package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		var temp []*TreeNode
		var part []int
		for len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			part = append(part, node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		res = append(res, part)
		queue = temp
	}
	return res
}
