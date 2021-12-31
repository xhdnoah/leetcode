package main

// Input: n = 3 Output: [[1,2,3],[8,9,4],[7,6,5]]
// 循环不变量原则 保持左开右闭不变
func generateMatrix(n int) [][]int {
	startx, starty := 0, 0 // 定义每循环一个圈的起始位置
	loop := n / 2          // 每个圈循环几次，若 n 为奇数 3 loop = 1 只是循环一圈
	mid := n / 2           // 矩阵中间位置 n 为 3 中间位置 (1，1) n 为 5 中间位置为 (2, 2)
	count := 1             // 给矩阵中的空格赋值
	offset := 1            // 控制循环中每一条边遍历的长度
	i, j := 0, 0
	res := make([][]int, n)
	for loop > 0 {
		i = startx
		j = starty

		// 模拟填充上行从左到右
		for j = starty; j < starty+n-offset; j++ {
			res[startx][j] = count
			count++
		}
		// 模拟填充右列从上到下(左闭右开)
		for i = startx; i < startx+n-offset; i++ {
			res[i][j] = count
			count++
		}
		// 模拟填充下行从右到左(左闭右开)
		for ; j > starty; j-- {
			res[i][j] = count
			count++
		}
		// 模拟填充左列从下到上(左闭右开)
		for ; i > startx; i-- {
			res[i][j] = count
			count++
		}

		// 第二圈开始时起始位置各自加 1 如：第一圈起始位置是(0, 0)，第二圈起始位置是(1, 1)
		startx++
		starty++

		// offset 控制每一圈里每条边遍历的长度
		offset += 2
	}

	// 如果 n 为奇数需单独给矩阵最中间的位置赋值
	if n % 2 {
		res[mid][mid] = count
	}
	return res
}
