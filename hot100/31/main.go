package main

func main() {
	link := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	//reverse(link)
	reverseKGroup(link, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
// https://www.bilibili.com/video/BV1kg4y137UU?vd_source=448ef99425176b9549d7f17447027da4
// 0 1 2 3 4
// 0 1 2
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, end := dummy, dummy
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		// 不足条件，达到末尾了
		if end == nil {
			break
		}
		// 保留后节点
		next := end.Next
		// 断开与后节点关系
		end.Next = nil
		// 开始翻转的节点
		first := pre.Next // 记录翻转的第一个节点
		// 取消绑定关系
		pre.Next = nil
		// 进行翻转，然后将末尾简单绑定到pre的Next
		pre.Next = reverse(first)
		//翻转之前的头节点绑定尾节点
		first.Next = next
		// 移动指针
		pre = first
		end = first
	}
	return dummy.Next
}

// 头插法，翻转链表
func reverse(link *ListNode) *ListNode {
	var pre *ListNode
	for link != nil {
		next := link.Next
		link.Next = pre
		pre = link
		link = next
	}
	return pre
}
