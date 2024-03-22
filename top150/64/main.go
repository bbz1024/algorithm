package main

import "fmt"

func main() {
	c := make([]int, 20)
	fmt.Println(c)
	fmt.Println(len(c[0:10]))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mp2 = make(map[int]int)

func build(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	tree := new(TreeNode)
	//根据前序创建节点
	tree.Val = preorder[0]
	Idx := 0
	for ; Idx < len(inorder); Idx++ {
		if inorder[Idx] == preorder[0] {
			break
		}
	}
	tree.Left = build(preorder[1:len(inorder[:Idx])+1], inorder[:Idx])
	tree.Right = build(preorder[len(inorder[:Idx])+1:], inorder[Idx+1:])
	return tree
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	for i := 0; i < len(inorder); i++ {
		mp2[inorder[i]] = i
	}
	return build(preorder, inorder)
}
