package main

func main() {
	copyRandomList2(&Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}})
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodes := make(map[*Node]*Node)
	for temp := head; temp != nil; temp = temp.Next {
		nodes[temp] = &Node{Val: temp.Val}
	}
	for temp := head; temp != nil; temp = temp.Next {
		nodes[temp].Next = nodes[temp.Next]
		nodes[temp].Random = nodes[temp.Random]
	}
	return nodes[head]
}
func copyRandomList2(head *Node) *Node {
	for temp := head; temp != nil; {
		node := &Node{Next: temp.Next, Val: temp.Val}
		temp.Next = node
		temp = node.Next
	}
	for temp := head; temp != nil; temp = temp.Next.Next {
		if temp.Random != nil {
			temp.Next.Random = temp.Random.Next

		}
	}
	cur := head.Next
	res := head.Next
	pre := head
	for cur.Next != nil {
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	pre.Next = nil
	return res
}
