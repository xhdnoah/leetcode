package main

import (
	"math"
	"unicode"
)

// Atoi 算法：读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// Input: s = "   -42" Output: -42 Input: s = "4193 with words" Output: 4193
// 有限状态自动机: 程序在每个时刻有一个状态 s，每次从序列中输入一个字符 c，并根据字符 c 转移到下一个状态 s'
// 这样只要建立一个覆盖所有情况的从 s 与 c 映射到 s' 的表格即可解决题目中的问题
// 四个状态：start 开始读入 end 结束读入 in_number(正在数字当中) signed 符号 +/-
//            ' '	  +/-	    number	  other
// start	    start	signed	in_number	end
// signed	    end	  end	    in_number	end
// in_number	end	  end	    in_number	end
// end	      end	  end	    end	      end
func myAtoi(s string) int {
	automation := Constructor()
	for _, c := range s {
		automation.get(c)
	}
	return automation.sign * automation.ans
}

type Automation struct {
	state string
	sign  int
	ans   int
	table map[string][]string
}

func Constructor() Automation {
	return Automation{
		state: "start",
		sign:  1,
		ans:   0,
		table: map[string][]string{
			"start":     {"start", "signed", "in_number", "end"},
			"signed":    {"end", "end", "in_number", "end"},
			"in_number": {"end", "end", "in_number", "end"},
			"end":       {"end", "end", "end", "end"},
		},
	}
}

func (auto *Automation) getCol(c rune) int {
	if c == ' ' {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if unicode.IsDigit(c) {
		return 2
	}
	return 3
}

func (this *Automation) get(c rune) {
	this.state = this.table[this.state][this.getCol(c)]
	if this.state == "in_number" {
		this.ans = this.ans*10 + int(c-'0')
		if this.sign == 1 {
			this.ans = min(this.ans, math.MaxInt32)
		} else {
			this.ans = min(this.ans, -math.MinInt32)
		}
	} else if this.state == "signed" {
		if c == '+' {
			this.sign = 1
		} else {
			this.sign = -1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
