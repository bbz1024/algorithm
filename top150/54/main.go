package main

import "math"

func main() {

}

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}
func (this *MinStack) Empty() bool {
	return len(this.stack) == 0
}
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() int {
	if this.Empty() {
		return 0
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if this.Empty() {
		return math.MaxInt32
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt32
	for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min
}
