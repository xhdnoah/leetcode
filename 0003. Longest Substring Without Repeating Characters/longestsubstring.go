package main

import . "leetcode/utils"

// Input: "pwwkew" Output: 3 ("wke")
// 滑动窗口：右边界不断右移，将未知字节合入窗口
// 再通过不断右移左边界保证新字节不重复，每次循环更新最值
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int) // 窗口内部字节出现频率
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		// 有点 K8s 控制器模型的思想
		// 我不关心 window[c] 具体值然后具体操作 
		// 只认准需要达到的状态 然后循环 reconcile 到目标状态
		for window[c] > 1 { 
			d := s[left]
			left++
			window[d]--
		}
		res = Max(res, right-left)
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
		res = Max(res, right-left+1)
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
		res = Max(res, right-left)
	}
	return res
}

func main() {
	lengthOfLongestSubstring("pwwkew")
}
