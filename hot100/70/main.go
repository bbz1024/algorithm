package main

func main() {
	//["MinStack","push","push","push","getMin","pop","top","getMin"]
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()
	minStack.Pop()
	minStack.Top()
	print(minStack.GetMin())
}

type MinStack struct {
	stack [][2]int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([][2]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, [2]int{val, val})
	} else {
		/*
			每次入栈时，进行判断入职元素是否小于当前栈顶元素维护的的最小值，这样一来可以进行动态的维护最小值
			因为能确保栈顶元素为未出去之前，后面的元素都是未出栈的元素
		*/
		this.stack = append(this.stack, [2]int{val, min(val, this.stack[len(this.stack)-1][1])})
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	/*
		题目要求GetMin()函数的时间复杂度为O(1)，所以需要再加入元素的时候维护最小值
	*/
	return this.stack[len(this.stack)-1][1]
}
