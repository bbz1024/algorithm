package main

import (
	"container/heap"
	"fmt"
)

func main() {
	heap.Init(&MaxHeap{})
	heap.Init(&MinHeap{})
	constructor := Constructor()
	constructor.AddNum(1)
	constructor.AddNum(2)
	constructor.AddNum(3)
	median := constructor.FindMedian()
	fmt.Println(median)
}

/* 大根堆元素比小根堆元素多1
heap1：大根对
heap2：小根对

当元素为偶数的时候往小跟对插入，把小根堆对顶元素插入到大根堆。中位数就是大根堆堆顶
当元素为奇数的时候插入到大跟对，把堆顶的元素一入到小根堆。中位数就是两个堆顶和/2

*/

// MaxHeap 是一个大顶堆的结构体。
type MaxHeap []int

// Len 返回堆中的元素数量。
func (h MaxHeap) Len() int { return len(h) }

// Less 比较两个元素，以确定它们的顺序。
// 对于大顶堆，i > j 表示 i 应该排在 j 前面。
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

// Less 比较两个元素，以确定它们的顺序。
// 对于小顶堆，i < j 表示 i 应该排在 j 前面。
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	min MinHeap
	max MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		min: MinHeap{},
		max: MaxHeap{},
	}

}
func (this *MedianFinder) AddNum(num int) {
	// 偶数
	if this.max.Len() == this.min.Len() {
		heap.Push(&this.min, num)
		heap.Push(&this.max, heap.Pop(&this.min))
	} else {
		heap.Push(&this.max, num)
		heap.Push(&this.min, heap.Pop(&this.max))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() == this.min.Len() {
		return float64((this.max)[0]+(this.min)[0]) / 2
	} else {
		return float64((this.max)[0])
	}
}

// -------------------- 超时 --------------------
