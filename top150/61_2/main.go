package main

import "fmt"

/*
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {
	nd1 := ListNode{Val: 1}
	nd2 := ListNode{Val: 2}
	nd1.Next = &nd2
	res := removeNthFromEnd(&nd1, 1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0, 31)
	dummy := &ListNode{
		Next: head,
	}
	nodes = append(nodes, dummy)
	cur := dummy.Next
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	l := len(nodes)
	start := l - n - 1
	if start >= l {
		nodes[start].Next = nil
	}
	nodes[start].Next = nodes[start].Next.Next
	return nodes[0].Next
}
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//双指针
	dummy1 := &ListNode{
		Next: head,
	}
	fast := dummy1
	slow := dummy1
	for n != 0 {
		n--
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy1.Next
}
