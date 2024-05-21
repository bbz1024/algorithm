package main

func main() {
	nd1 := ListNode{Val: 1}
	nd2 := ListNode{Val: 3}
	nd3 := ListNode{Val: 3}
	nd4 := ListNode{Val: 2}
	nd5 := ListNode{Val: 5}
	nd6 := ListNode{Val: 2}
	nd1.Next = &nd2
	nd2.Next = &nd3
	nd3.Next = &nd4
	nd4.Next = &nd5
	nd5.Next = &nd6
	partition(&nd1, 3)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	smallDummy := &ListNode{}
	lagerDummy := &ListNode{}
	small := smallDummy
	lager := lagerDummy
	for head != nil {
		if head.Val >= x {
			lager.Next = head
			lager = lager.Next
		} else {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}
	//TODO 为什么需要lager.Next=nil
	lager.Next = nil
	small.Next = lagerDummy.Next

	return smallDummy.Next
}
