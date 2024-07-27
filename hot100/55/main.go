package main

func main() {

}

/*
有递归就一定存在回溯
*/
func permute(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	var permution func(level int)
	permution = func(level int) {
		if level == length {
			res = append(res, append([]int(nil), nums...))
			return
		}
		for i := level; i < length; i++ {
			// 交换
			nums[level], nums[i] = nums[i], nums[level]
			permution(level + 1)
			// 回溯(撤消)
			nums[level], nums[i] = nums[i], nums[level]
		}
	}
	permution(0)
	return res
}
func permute2(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	var used = make([]bool, length)
	var path = make([]int, length)
	var dfs func(l int)
	dfs = func(l int) {
		if l == length {
			res = append(res, append([]int(nil), path...))
			return
		}
		for i, b := range used {
			if !b {
				used[i] = true
				path[l] = nums[i]
				dfs(l + 1)
				used[i] = false
			}
		}
	}
	dfs(0)
	return res
}

// 回溯模板
/*
func backtrack(path, choices []int) {
	if len(path) == len(choices) {
		// 终止条件,收取结果
		res = append(res, append([]int(nil), path...))
		return
	}
	for i := 0; i < len(choices); i++ {
		// 选择
		path = append(path, choices[i])
		// 递归
		backtrack(path, choices) // 通常在递归之后开始回溯
		// 回溯（撤消）
		path = path[:len(path)-1] // 撤销 选择
}
*/
