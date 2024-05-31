package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	dummy := &ListNode{
		Next: headA,
	}
	cur := dummy.Next
	for ; cur != nil; cur = cur.Next {
		mp[cur] = struct{}{}

	}
	dummy.Next = headB
	cur = dummy.Next
	for ; cur != nil; cur = cur.Next {
		if _, ok := mp[cur]; ok {
			return cur
		}
	}
	return nil
}

// 如果两个链表相交，那么相交点之后的长度是相同的
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != nil || p2 != nil {
		if p1 == p2 {
			return p1
		}
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return nil
}
