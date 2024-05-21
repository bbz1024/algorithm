package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var answer [][]int
	for len(queue) != 0 {
		length := len(queue)
		var c []int
		for i := 0; i < length; i++ {
			node := queue[0]
			c = append(c, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		answer = append(answer, c)
	}
	return answer
}
