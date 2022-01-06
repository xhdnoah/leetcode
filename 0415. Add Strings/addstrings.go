package main

import "strconv"

// Input: num1 = "456", num2 = "77" Output: "533"
func addStrings(num1 string, num2 string) string {
	res, i, j, carry := "", len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		} else {
			n2 = 0
		}
		tmp := n1 + n2 + carry
		carry = tmp / 10
		res = strconv.Itoa(tmp%10) + res
		i, j = i-1, j-1
	}
	if carry > 0 {
		return "1" + res
	}
	return res
}
