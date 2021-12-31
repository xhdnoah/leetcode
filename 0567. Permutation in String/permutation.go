package main

import "fmt"

// Input: s1 = "ab", s2 = "eidbaooo" Output: true
// Input: s1 = "ab", s2 = "eidboaoo" Output: false
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s2) {
		c := s2[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left == len(s1) { // 判断左侧是否收缩, 将窗口控制在 len(s1)-1
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func main() {
	shouldTrue := checkInclusion("ab", "eidbaooo")
	shouldFalse := checkInclusion("ab", "eidboaoo")
	fmt.Println(shouldTrue, shouldFalse)
}
