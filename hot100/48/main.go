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

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	var inOrder func(root *TreeNode, sum int)
	inOrder = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		if sum+root.Val == targetSum {
			res++
		}
		inOrder(root.Left, sum+root.Val)
		inOrder(root.Right, sum+root.Val)
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		inOrder(node, 0)
		stack = append(stack, node.Left, node.Right)
	}
	return res
}
