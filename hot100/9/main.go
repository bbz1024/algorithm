package main

import "fmt"

func main() {
	res := findAnagrams("cbaebabacd", "abc")
	fmt.Println(res)
}
func findAnagrams(s, p string) (ans []int) {
	sl, pl := len(s), len(p)
	if pl > sl {
		return
	}
	var pArr [26]byte
	var sArr [26]byte
	for i := 0; i < pl; i++ {
		pArr[p[i]-'a']++
	}

	for l, r := 0, 0; r < sl; r++ {
		sArr[s[r]-'a']++
		if r-l+1 > pl {
			sArr[s[l]-'a']--
			l++
		}
		if r-l+1 == pl {
			if pArr == sArr {
				ans = append(ans, l)
			}
		}
	}
	return
}
