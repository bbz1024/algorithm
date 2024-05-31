package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	mp := make(map[*ListNode]struct{})
	p := head
	for p != nil {
		if _, ok := mp[p.Next]; ok {
			return p.Next
		}
		mp[p] = struct{}{}
		p = p.Next
	}
	return nil
}
