package main

import "fmt"

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	res := reverseBetween(node1, 2, 4)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
