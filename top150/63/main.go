package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Compare(node *TreeNode, node2 *TreeNode) bool {
	//如果不存在左右节点返回true
	if node == nil && node2 == nil {
		return true
	}
	//left或right都为nil,或者left和right有一个为nil
	if node == nil || node2 == nil {
		return false
	}
	//值不相等
	if node.Val != node2.Val {
		return false
	}
	//递归比较
	return Compare(node.Left, node2.Right) && Compare(node.Right, node2.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Compare(root.Left, root.Right)

}
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	if root.Right == nil && root.Left == nil {
		return true
	}
	if root.Right == nil || root.Left == nil {
		return false
	}
	queue = append(queue, root.Left, root.Right)

	for len(queue) != 0 {
		node := queue[0]
		node2 := queue[1]
		queue = queue[2:]
		if node == nil && node2 == nil {
			continue
		}
		//left或right都为nil,或者left和right有一个为nil
		if node == nil || node2 == nil {
			return false
		}
		if node.Val != node2.Val {
			return false
		}
		queue = append(queue, node.Left, node2.Right, node.Right, node2.Left)

	}
	return true
}
