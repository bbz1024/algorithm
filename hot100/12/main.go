package main

func main() {
}

func minWindow(s string, t string) string {
	var isCovered = func(cntS, cntT []int) bool {
		for i := 'A'; i <= 'Z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		for i := 'a'; i <= 'z'; i++ {
			if cntS[i] < cntT[i] {
				return false
			}
		}
		return true
	}
	m := len(s)
	ansLeft, ansRight, left := -1, m, 0
	var cntS, cntT [128]int
	for i := 0; i < len(t); i++ {
		cntT[t[i]]++
	}
	for right, v := range s { // 移动子串右端点
		cntS[v]++ // 右端点字母移入子串
		//存在覆盖
		for isCovered(cntS[:], cntT[:]) {
			//找到更短的子串
			if ansRight-ansLeft > right-left {
				ansRight, ansLeft = right, left
			}
			//左端点字母移除子串
			cntS[s[left]]--
			//移动子串左端点
			left++
		}
	}
	if ansLeft < 0 {
		return ""
	}
	return s[ansLeft : ansRight+1]
}

/*
func minWindow(s string, t string) string {
	if s == t {
		return t
	}
	if len(s) < len(t) {
		return ""
	}
	var sMap [52]int
	var tMap [52]int
	for i := 0; i < len(t); i++ {
		tMap[ToNum(t[i])]++
	}
	var ans string
	l, r := 0, 0
	for r < len(s) {
		if tMap[ToNum(s[r])] > 0 {
			sMap[ToNum(s[r])]++
		}
		//大于了，进行跳转到下一个
		//	移动l
		if tMap[ToNum(s[r])] < sMap[ToNum(s[r])] {
			sMap[ToNum(s[r])]--
			l++
			for l<r&&sMap[ToNum(s[l])] == 0 {
				l++
			}
		}
		//收货一次结果
		if tMap == sMap {
			if ans != "" {
				if len(ans) > r-l+1 {
					ans = s[l : r+1]
				}
			} else {
				ans = s[l : r+1]
			}
			if tMap[ToNum(s[l])] > 0 {
				sMap[ToNum(s[l])]--
			}
			l++
			for l < r && tMap[ToNum(s[l])] == 0 {
				l++
			}
		}
		r++
	}
	return ans
}
func ToNum(v byte) int {
	if v >= 'a' && v <= 'z' {
		return int(v-'a') + 26
	}
	return int(v - 'A')
}
*/
