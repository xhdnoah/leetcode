package main

// Input: s = "barfoothefoobarman", words = ["foo","bar"] Output: [0,9]
// 先将数组里所有元素存到 map 中累计出现次数，再从源字符串从头开始扫，每次判断数组元素是否全都用完了(计数为 0)
// 如果用完且长度正好为任意排列组合总长，则记录该组合的起始下标；如不符合继续考察源字符串下一个字符，直到扫完整个源字符串
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}

	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	// start 为当前窗口的起点，i 为单个 word 起点，注意边界条件 i + length < len(s) + 1 假设最后一个 word 首字母 index + length 的值即为 len(s)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		if tmpCounter[s[i:i+length]] > 0 { // s 在[i, i+length] 上的片段存在于 words 数组
			tmpCounter[s[i:i+length]]--
			if checkWords(tmpCounter) && (i+length-start == totalLen) { // words 用尽且子串长度为数组元素排列总长
				res = append(res, start)
				continue
			}
			i = i + length - 1 // 注意减一，这里满足一个 word 后跳转到下一个 word 起点开始扫描
		} else {
			start++
			i = start - 1                 // 不满足 word 后，起点跳转到整个窗口的 start
			tmpCounter = copyMap(counter) // 重新开始操作计数器
		}
	}
	return res
}

// 判断数组元素是否用完
func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}
