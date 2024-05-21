package main

/*
https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/?envType=study-plan-v2&envId=top-interview-150
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head, Val: -1010}
	slow := dummy
	fast := dummy
	for fast != nil && fast.Next != nil {
		if fast.Val == fast.Next.Val {
			for slow.Next.Val != fast.Val {
				slow = slow.Next
			}
			for fast.Next != nil && fast.Val == fast.Next.Val {
				fast = fast.Next
			}
			slow.Next = fast.Next
		}
		fast = fast.Next
	}
	return dummy.Next
}
