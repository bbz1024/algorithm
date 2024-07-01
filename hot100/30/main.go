package main

func main() {
	swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	})

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	temp := dummy
	for temp.Next != nil && temp.Next.Next != nil {
		next1 := temp.Next      // 1
		next2 := temp.Next.Next // 2
		temp.Next = next2
		next1.Next = next2.Next
		next2.Next = next1

		temp = next1
	}
	return dummy.Next
}
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// head 连接后面交换完成的子链表，next连接head，完成交换
	next := head.Next
	head.Next = swapPairs2(next)
	next.Next = head
	return next
}
