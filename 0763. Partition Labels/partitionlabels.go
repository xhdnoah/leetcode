package main

// Input: S = "ababcbacadefegdehijhklij" Output: [9,7,8]
// Explanation: The partition is "ababcbaca", "defegde", "hijhklij".
// This is a partition so that each letter appears in at most one part.
// 给出一个字符串，要求输出满足条件窗口的长度，条件是在这个窗口内，字母中出现在这一个窗口内，不出现在其他窗口内。
// 第一种思路先记录每个字母出现次数，再对滑动窗口中各字母判断次数是否用尽，如果窗口内所有字母次数都为 0，即符合条件
func partitionLabels_counter(S string) []int {
	visit, counter, res, sum, lastLength := make([]int, 26), map[byte]int{}, []int{}, 0, 0
	for i := 0; i < len(S); i++ {
		counter[S[i]]++
	}

	for i := 0; i < len(S); i++ {
		counter[S[i]]--
		visit[S[i]-'a'] = 1 // 记录已遍历范围内出现的字符
		sum = 0
		for j := 0; j < 26; j++ {
			if visit[j] == 1 {
				sum += counter[byte('a'+j)] // 所有已出现字符的计数器均为 0
			}
		}
		if sum == 0 {
			res = append(res, i+1-lastLength)
			lastLength = i + 1
		}
	}
	return res
}

// 另一种思路记录每个字符最后一次出现的下标，在每个滑动窗口中，依次判断每个字母最后一次出现的位置
// 如果在一个下标内，所有字母的最后一次出现的位置都包含进来了，那么这个下标就是这个满足条件的窗口大小
func partitionLabels_index(S string) []int {
	var lastIndexOf [26]int
	for i, v := range S {
		lastIndexOf[v-'a'] = i
	}
	var res []int
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = lastIndexOf[S[start]-'a'] // end 为窗口首个字符最后出现的位置
		for i := start; i < end; i++ {
			if end < lastIndexOf[S[i]-'a'] { // 如果窗口中间存在最后位置超出 end 的情况
				end = lastIndexOf[S[i]-'a'] // 则 end 的位置跳转至此
			}
		}
		res = append(res, end-start+1)
	}
	return res
}
