package main

import (
	"sort"
	"strconv"
	"strings"
)

type atom struct {
	name  string
	count int
}

type atoms []atom

// Input: formula = "K4(ON(SO3)2)2" Output: "K4N2O14S4"
// The count of elements are {'K': 4, 'N': 2, 'O': 14, 'S': 4}.
// 利用栈处理每个化学元素，用 map 记录每个化学元素的个数，注意字母大小写问题

func (this atoms) Len() int           { return len(this) }
func (this atoms) Less(i, j int) bool { return strings.Compare(this[i].name, this[j].name) < 0 }
func (this atoms) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this atoms) String() string {
	s := ""
	for _, a := range this {
		s += a.name
		if a.count > 1 {
			s += strconv.Itoa(a.count)
		}
	}
	return s
}

func countOfAtoms(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	stack := make([]string, 0) // 元素栈
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' || c == ')' { // 括号入栈
			stack = append(stack, string(c))
		} else if isUpperLetter(c) { // 大写是新元素开始
			j := i + 1
			for ; j < n; j++ { // 查找下一个非小写字母以定位新元素
				if !isLowerLetter(s[j]) {
					break
				}
			}
			stack = append(stack, string(s[i:j]))
			i = j - 1
		} else if isDigital(c) {
			j := i + 1
			for ; j < n; j++ {
				if !isDigital(s[j]) {
					break
				}
			}
			stack = append(stack, s[i:j])
			i = j - 1
		}
	}

	count, deep := make([]map[string]int, 100), 0
	for i := 0; i < 100; i++ {
		count[i] = make(map[string]int)
	}
	// s: "Mg(OH)2" stack: ["Mg" "(" "O" "H" ")" "2"]
	for i := 0; i < len(stack); i++ {
		t := stack[i]
		if isUpperLetter(t[0]) { // 新元素
			num := 1
			if i+1 < len(stack) && isDigital(stack[i+1][0]) {
				num, _ = strconv.Atoi(stack[i+1])
				i++
			}
			count[deep][t] += num // deep 记录括号深度
		} else if t == "(" {
			deep++
		} else if t == ")" {
			num := 1
			if i+1 < len(stack) && isDigital(stack[i+1][0]) {
				num, _ = strconv.Atoi(stack[i+1]) // 括号后的数字
				i++
			}
			for k, v := range count[deep] {
				count[deep-1][k] += v * num // 深度减一前对其中元素在上个深度数组中的个数进行累加
			}
			count[deep] = make(map[string]int)
			deep--
		}
	}
	as := atoms{}
	for k, v := range count[0] {
		as = append(as, atom{name: k, count: v})
	}
	sort.Sort(as)
	return as.String()
}

func isDigital(v byte) bool {
	if v >= '0' && v <= '9' {
		return true
	}
	return false
}

func isUpperLetter(v byte) bool {
	if v >= 'A' && v <= 'Z' {
		return true
	}
	return false
}

func isLowerLetter(v byte) bool {
	if v >= 'a' && v <= 'z' {
		return true
	}
	return false
}

func main() {
	countOfAtoms("Mg(OH)2")
}
