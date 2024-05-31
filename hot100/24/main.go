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
	isPalindrome(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转比较
// 栈
func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{}
	cur := dummy
	p := head
	for ; p != nil; p = p.Next {
		node := &ListNode{
			Val: p.Val,
		}
		node.Next = cur.Next
		cur.Next = node

	}
	cur = dummy.Next
	for cur.Val == head.Val {
		if cur.Next == nil {
			return true
		}
		cur = cur.Next
		head = head.Next
	}
	return false
}

func isPalindrome2(head *ListNode) bool {
	var f func(head *ListNode) bool
	frontPointer := head
	f = func(node *ListNode) bool {
		if node != nil {
			// 这里先遍历到head的最尾部，
			if !f(node.Next) {
				return false
			}
			if node.Val != frontPointer.Val {
				return false
			}
			//开始重递归栈里出结果，与其进行比较
			frontPointer = frontPointer.Next
		}
		return false
	}
	return f(head)
}
