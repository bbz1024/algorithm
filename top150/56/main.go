package main

import (
	"fmt"
	"math"
)

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
	fmt.Println(addTwoNumbers(&nd1, &nd1))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func calList(li *ListNode, answer, level int) int {
	if li == nil {
		return answer
	}
	answer = answer + li.Val*int(math.Pow(10, float64(level)))
	return calList(li.Next, answer, level+1)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := calList(l1, 0, 0) + calList(l2, 0, 0)
	dummy := &ListNode{
		Next: nil,
	}
	c := dummy
	for res != 0 {
		nd := &ListNode{
			Val: res % 10,
		}
		c.Next = nd
		c = nd
		res /= 10

	}
	return dummy.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Next: nil}
	carry := 0
	p := l1
	q := l2
	curr := dummy

	for p != nil || q != nil {
		x := 0
		y := 0
		if p != nil {
			x = p.Val
			p = p.Next
		}
		if q != nil {
			y = q.Val
			q = q.Next
		}

		sum := x + y + carry
		curr.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
