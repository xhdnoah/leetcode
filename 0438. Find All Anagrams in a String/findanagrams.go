package main

// Input: s = "cbaebabacd", p = "abc" Output: [0,6]
// Input: s = "abab", p = "ab" Output: [0,1,2]
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right, valid, res := 0, 0, 0, []int{}
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left == len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

func main() {
	// res := findAnagrams("cbaebabacd", "abc")
	// res := findAnagrams("abab", "ab")
}
