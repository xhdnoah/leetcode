package main

// Input: S = "ADOBECODEBANC", T = "ABC" Output: "BANC"
// 滑动窗口，在滑动过程中不断包含 T，直到完全包含 T 的字符后，记下左右窗口位置和窗口大小，每次更新符合条件的窗口和窗口大小的最小值
// 整个过程窗口先从起点向右开始增长, count 到头后左边界向右收缩
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	// tFreq: T 串所有字符出现频率; sFreq: S 中窗口内各字符出现频率
	var tFreq, sFreq [256]int
	// count 为窗口中包含 T 中字符的个数
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			// 右移前检查窗口最右侧字符剩余次数 <= tFreq 中对应字符次数，说明该字符属于 T 且窗口中该字符个数还未超出上限 count++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else { // 满足 else 表示 count 或 right 到头
			if right-left+1 < minW && count == len(t) { // 出现更小的窗口长度
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			// 窗口左移前检查最左侧字符，频率相等说明存在于 T 则 count-- 如果是 > 关系说明不存在于 T 这里不会是 <
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}
