package main

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node1.Random = node2
	node3.Random = node3
	node4.Random = node2
	copyRandomList(node1)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	dummy := &Node{Next: head}
	cur := dummy.Next
	for ; cur != nil; cur = cur.Next {
		nd := &Node{
			Val: cur.Val,
		}
		mp[cur] = nd

	}
	cur = dummy.Next
	for ; cur != nil; cur = cur.Next {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
	}
	return mp[head]
}
