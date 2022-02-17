package main

// Given the API rand7() that generates a uniform random integer in the range [1, 7]
// write a function rand10() that generates a uniform random integer in the range [1, 10]
// 主要理解二进制的每一位出现 0和 1的可能性都是相同的，所以能生成的每一个数概率都是相同的 10 的二进制为 1010
func rand10() int {
	ans := rand2()
	for i := 0; i < 3; i++ {
		ans <<= 1      // 左移一位
		ans ^= rand2() // 按位异或
	}
	if ans <= 10 && ans > 0 {
		return ans
	} else {
		return rand10()
	}
}

func rand2() int {
	ans := rand7()
	if ans == 7 {
		return rand2()
	} else {
		return ans % 2
	}
}
