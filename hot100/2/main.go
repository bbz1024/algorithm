package main

import "sort"

func main() {

}

// 排序、哈希
func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	mp := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		t := []byte(strs[i])
		sort.Slice(t, func(j, k int) bool {
			return t[j] > t[k]
		})
		v, ok := mp[string(t)]
		if !ok {
			mp[string(t)] = len(ans)
			ans = append(ans, []string{strs[i]})
			continue
		}
		ans[v] = append(ans[v], strs[i])
	}
	return ans
}
