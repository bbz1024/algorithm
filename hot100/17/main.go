package main

func main() {

}

/*
https://leetcode.cn/problems/first-missing-positive/solutions/304743/que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution/?envType=study-plan-v2&envId=top-100-liked
*/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	// 获取未出现最小正整数，那么正整数访问肯定在 [1,len(num)+1] 间
	// 1. 进行标记负数的为 n+ 1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	//2. 将小于n对应的位置变为负数 改为负号就是进行标记已经出现过了
	for i := 0; i < n; i++ {
		num := abs(nums[i]) // 重复的例如[4,1,1,3,2]  这里的两个1会把4修改为-4再到4，这样的话会出现错误，非负的话就说明1是不存在的
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	//返回第一个大于0的下标+1，就是最小正数的值
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
