package main

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums = []int{1, 2, 1, 1, 1}
	res := jump(nums)
	println(res)
}
func jump(nums []int) int {
	var cur, next, res int
	/*
		cur：当前覆盖的最远距离下标
		next：记录走的最大步数
		res：走的步数
	*/
	l := len(nums)
	if l <= 1 {
		return 0
	}

	for i := 0; i < l; i++ {
		next = max(next, i+nums[i]) // 更新下一步能达到的最远距离
		// 走到了覆盖范围
		if cur == i { // 遇到当前覆盖的最远距离下标
			if cur != l-1 {
				// 步数往前移
				res++
				cur = next // 到最大覆盖范围
			} else {
				return res
			}
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
