package main

// Input: "pwwkew" Output: 3 ("wke")
// 滑动窗口：右边界不断右移，只要没有重复字符，就持续向右扩大窗口边界
// 出现重复字符缩小左边界，直到重复字符移出左边界，再继续移动右边界，每次移动更新最大长度
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)
	}
	return res
}

// 优化了空间使用 通过存储字节 x - 'a' 的差值来表示 x 出现在滑动窗口中的频率
func lengthOfLongestSubstring_freq(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [256]int
	res, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left--
		}
		res = max(res, right-left+1)
	}
	return res
}

// 滑动窗口-哈希桶 map 的键值分别为字节和其在字符串中的下标
func lengthOfLongestSubstring_map(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for right < len(s) {
		if idx, ok := indexes[s[right]]; ok && idx >= left {
			// 当 right 撞到重复字节，left 跳至前一个该字节的后面一位
			left = idx + 1
		}
		indexes[s[right]] = right
		right++
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	lengthOfLongestSubstring("pwwkew")
}
