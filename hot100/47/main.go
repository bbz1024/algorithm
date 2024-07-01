package main

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(tree)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先序遍历
func flatten(root *TreeNode) {
	var preOrder func(root *TreeNode)
	var nodes []*TreeNode
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes = append(nodes, root)
		preOrder(root.Left)
		root.Left = nil
		preOrder(root.Right)

	}
	preOrder(root)
	// 先序遍历，获取顺序，重新构建一次树
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Right = nodes[i]
	}
}
