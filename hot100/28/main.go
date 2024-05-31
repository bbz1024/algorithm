package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	}
	fmt.Println(l1, l2)
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func cal(l1 *ListNode, level int) int {
	if l1 == nil {
		return 0
	}
	t := 1
	for i := 0; i < level; i++ {
		t *= 10
	}
	return l1.Val*t + cal(l1.Next, level+1)
}

// 存在进位丢失
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := cal(l1, 0) + cal(l2, 0)
	if sum == 0 {
		return &ListNode{Val: 0}
	}
	dummy := &ListNode{}
	cur := dummy
	for sum != 0 {
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		sum /= 10
	}
	return dummy.Next
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 初始化进位值
	carry := 0

	// 创建哑节点
	dummy := &ListNode{}
	cur := dummy

	// 循环直到没有进位
	for l1 != nil || l2 != nil || carry != 0 {
		// 如果链表 l1 或 l2 还有剩余节点，或者当前有进位
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		} else {
			val1 = 0
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		} else {
			val2 = 0
		}

		// 计算当前位的和加上进位
		total := val1 + val2 + carry

		// 更新当前节点的值和进位
		carry = total / 10
		cur.Next = &ListNode{Val: total % 10}
		cur = cur.Next
	}

	// 返回新链表的头部
	return dummy.Next
}
