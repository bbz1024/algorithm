package main

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Right.Right = &Node{Val: 5}
	root.Right.Right.Left = &Node{Val: 6}
	root.Right.Right.Right = &Node{Val: 7}
	root.Right.Right.Right.Right = &Node{Val: 8}
	connect(root)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue = []*Node{root}
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if i == 0 {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
				continue
			}
			queue[i-1].Next = queue[i]
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return root
}

/*
     1
 2		3
4		  5



*/
