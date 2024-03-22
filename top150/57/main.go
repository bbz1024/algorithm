package main

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

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Next: nil,
	}
	cur := dummy
	for list1 != nil && list2 != nil {
		nd := &ListNode{}
		if list1.Val > list2.Val {
			nd.Val = list2.Val
			list2 = list2.Next
		} else {
			nd.Val = list1.Val
			list1 = list1.Next
		}
		cur.Next = nd
		cur = nd
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
