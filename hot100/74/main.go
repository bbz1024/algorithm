package main

func main() {

}
func findKthLargest(nums []int, k int) int {
	bigHeap := NewBigHeap(nums)
	return bigHeap.TopK(k)[k-1]
}

type BigHeap struct {
	arr []int
}

func NewBigHeap(arr []int) *BigHeap {
	b := &BigHeap{
		arr: arr,
	}
	b.construct()
	return b
}
func (b *BigHeap) construct() {
	/*
		找到最后一个非叶子节点,调整该节点，然后再调整兄弟节点

		左右孩子节点：2i+1 2i+2
		父节点：(i-1)/2
		8,9,1,2,3,6,5,7,4
						8
				9				1
			2		3		6		5
		  7	  4
		最后一个非叶子节点：
		(9-1)/2=4 => 3

						9
				8				6
			7		3		1		5
		  4	  2

	*/
	for i := (len(b.arr) - 1) / 2; i >= 0; i-- {
		b.adjust(i)
	}
}
func (b *BigHeap) TopK(k int) []int {
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, b.arr[0])
		b.arr[0], b.arr[len(b.arr)-1] = b.arr[len(b.arr)-1], b.arr[0]
		b.arr = b.arr[:len(b.arr)-1]
		b.adjust(0)
	}
	return res
}
func (b *BigHeap) adjust(i int) {
	left := 2*i + 1
	right := 2*i + 2
	if left >= len(b.arr) {
		return
	}
	//	只有一个左
	if right >= len(b.arr) {
		if b.arr[i] < b.arr[left] {
			b.arr[i], b.arr[left] = b.arr[left], b.arr[i]
			b.adjust(left)
		}
		return

	}
	// 有两个
	if b.arr[i] < b.arr[left] || b.arr[i] < b.arr[right] {
		if b.arr[left] > b.arr[right] {
			b.arr[i], b.arr[left] = b.arr[left], b.arr[i]
			b.adjust(left) // 重新调整子树
		} else {
			b.arr[i], b.arr[right] = b.arr[right], b.arr[i]
			b.adjust(right)
		}
	}
}
