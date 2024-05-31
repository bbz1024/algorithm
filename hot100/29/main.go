package main

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	removeNthFromEnd(list, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy
	for n != 0 {
		n--
		fast = fast.Next
	}
	// 快指针走到尾了，就说明slow到达了要删除的节点的前一个
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
