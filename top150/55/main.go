package main

import "fmt"

//https://leetcode.cn/problems/linked-list-cycle/?envType=study-plan-v2&envId=top-interview-150
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func main() {
	nd1 := ListNode{Val: 1}
	nd2 := ListNode{Val: 2}
	nd3 := ListNode{Val: 3}
	nd1.Next = &nd2
	nd2.Next = &nd3
	nd3.Next = &nd1
	fmt.Println(hasCycle(&nd1))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	/*
		// 1. 快慢指针
		if head == nil {
			return false
		}
		slow := head
		fast := head
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
		return false

	*/
	// 2.哈希表
	mp := make(map[*ListNode]struct{})
	dummy := &ListNode{
		Next: head,
	}
	nd := dummy.Next
	for nd != nil {
		if _, ok := mp[nd]; ok {
			return true
		}
		mp[nd] = struct{}{}
		nd = nd.Next
	}
	return false
}
