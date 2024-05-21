package main

import "fmt"

func main() {
	nd1 := ListNode{Val: 1}
	nd2 := ListNode{Val: 2}
	nd3 := ListNode{Val: 3}
	nd4 := ListNode{Val: 4}
	nd5 := ListNode{Val: 5}
	nd1.Next = &nd2
	nd2.Next = &nd3
	nd3.Next = &nd4
	nd4.Next = &nd5
	res := rotateRight(&nd1, 2)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1 // 链表大小
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	//环型
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
