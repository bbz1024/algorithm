package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var depth func(root *TreeNode, level int)
	depth = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(res) == level { // 这个深度首次遇到
			res = append(res, root.Val)

		}
		depth(root.Right, level+1) // 先递归右子树，保证首次遇到的一定是最右边的节点
		depth(root.Left, level+1)
	}
	depth(root, 0)
	return res
}
