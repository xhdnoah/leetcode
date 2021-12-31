package main

// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202" Output: 6
// A sequence of valid moves would be "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202".
// Note that a sequence like "0000" -> "0001" -> "0002" -> "0102" -> "0202" would be invalid,
// because the wheels of the lock become stuck after the display becomes the dead end "0102".
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	visited := map[string]bool{"0000": true}
	for _, s := range deadends {
		if s == "0000" {
			return -1
		}
		visited[s] = true
	}
	step := 0
	q := []string{"0000"}

	for len(q) > 0 {
		sz := len(q)
		// 遍历当前队列中的结点
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			if cur == target {
				return step
			}

			// 将一个结点的未遍历相邻结点加入队列
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if !visited[up] {
					q = append(q, up)
					visited[up] = true
				}
				down := minusOne(cur, j)
				if !visited[down] {
					q = append(q, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}

// 向上拨动一次
func plusOne(s string, j int) string {
	chars := []rune(s) // '0000' -> [48, 48, 48, 48]
	if chars[j] == '9' {
		chars[j] = '0'
	} else {
		chars[j] += 1
	}
	return string(chars)
}

// 向下拨动一次
func minusOne(s string, j int) string {
	chars := []rune(s)
	if chars[j] == '0' {
		chars[j] = '9'
	} else {
		chars[j] -= 1
	}
	return string(chars)
}
