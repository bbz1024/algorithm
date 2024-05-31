package main

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针 ， hash表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p := &ListNode{Next: head}

	slow := p.Next
	fast := p.Next
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
