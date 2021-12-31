package main

// anagram: 同字母异序词
// Input: s = "anagram", t = "nagaram" Output: true
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if freq[t[i]-'a'] == 0 {
			return false
		}
		freq[t[i]-'a']--
	}
	return true
}
