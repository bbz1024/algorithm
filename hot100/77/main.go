package main

func main() {

}
func maxProfit(prices []int) int {
	var res int
	var buy int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[buy] {
			res = max(res, prices[i]-prices[buy])
		} else {
			buy = i
		}
	}
	return res
}

// -------------------- 超时 --------------------
func maxProfit2(prices []int) int {
	var res int
	//暴力
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
