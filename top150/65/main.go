package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	val := postorder[len(postorder)-1]
	node := &TreeNode{
		Val: val,
	}
	i := 0
	for ; i < len(postorder); i++ {
		if inorder[i] == val {
			break
		}
	}
	//根据i找到左边部分：inorder[:len(postorder[:i])] ，后遍历的左： postorder[:i]
	node.Left = buildTree(inorder[:len(postorder[:i])], postorder[:i])
	//根据i找到右边部分：inorder[i+1:  后序的右：postorder[i:len(postorder)-1]
	node.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return node
}
