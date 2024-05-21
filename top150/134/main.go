package main

func main() {
	//^5 + 1 =》-5
	//^-5 + 1=>5
	myPow(2, 10)
}
func makePositive(n int) int {
	if n < 0 {
		return -n // 如果数值为负，返回其相反数（即绝对值）
	}
	return n // 数值为正或零，直接返回原数值
}

// 超出时间限制
func myPow(x float64, n int) float64 {
	var answer float64 = 1
	for i := 0; i < makePositive(n); i++ {
		answer *= x
	}
	if n < 0 {
		return 1 / answer
	}
	if n == 0 {
		return 1
	}
	return answer
}
func myPow2(x float64, n int) float64 {

}
