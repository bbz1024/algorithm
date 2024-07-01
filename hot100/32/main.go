package main

import (
	"math"
)

func main() {
	sortList(&ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{4, nil}}}}})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	hd := &ListNode{Val: math.MinInt32}
	p := head
	for p != nil {
		temp := p.Next
		q := hd
		for q.Next != nil && q.Next.Val < p.Val {
			q = q.Next
		}
		p.Next = q.Next
		q.Next = p
		p = temp
	}
	return hd.Next

}

/*
func sortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead.Next
	for cur != nil && cur.Next != nil {
		if cur.Val > cur.Next.Val {
			cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
			cur = dummyHead.Next
		} else {

			cur = cur.Next
		}

	}
	return dummyHead.Next
}


*/
