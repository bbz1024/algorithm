package main

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	reverseList2(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头插法
// 栈
func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for ; head != nil; head = head.Next {
		node := &ListNode{
			Val: head.Val,
		}
		node.Next = cur.Next
		cur.Next = node

	}
	return dummy.Next
}
func reverseList2(head *ListNode) *ListNode {
	var last *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}
