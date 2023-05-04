package main

import (
	"bytes"
	"sort"
)

// 使用哈希表记录每个字符出现的频率，将字符去重后存入列表，再将列表中的字符按照频率降序排序。
func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}
	type pair struct {
		ch  byte
		cnt int
	}
	pairs := make([]pair, 0, len(cnt))
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })
	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}
