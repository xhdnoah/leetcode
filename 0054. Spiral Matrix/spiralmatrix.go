package main

const (
	right = iota
	down
	left
	up
)

var dirs = []int{right, down, left, up}

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]] Output: [1,2,3,6,9,8,7,4,5]
// 通过判断左右上下边界有没有交叉判断终止。每次扫完一条边缩减对应的边界，e.g. 从左往右扫完，那么上边界就该-1
func spiralOrder(matrix [][]int) (res []int) {
	m, n := len(matrix), len(matrix[0])
	res = make([]int, 0, m*n)

	l, r, top, bottom := 0, n-1, 0, m-1            // 分别表示左右上下边界
	for cnt := 0; l <= r && top <= bottom; cnt++ { // 边界交叉则调整方向
		switch dirs[cnt%len(dirs)] {
		case right:
			for i, j := top, l; j <= r; j++ {
				res = append(res, matrix[i][j])
			}
			top++
		case down:
			for i, j := top, r; i <= bottom; i++ {
				res = append(res, matrix[i][j])
			}
			r--
		case left:
			for i, j := bottom, r; j >= l; j-- {
				res = append(res, matrix[i][j])
			}
			bottom--
		case up:
			for i, j := bottom, l; i >= top; i-- {
				res = append(res, matrix[i][j])
			}
			l++
		}
	}
	return
}
