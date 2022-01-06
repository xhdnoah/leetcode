package main

// Input: n = 3 Output: [[1,2,3],[8,9,4],[7,6,5]]
func generateMatrix(n int) [][]int {
	l, r, t, b := 0, n-1, 0, n-1 // 代表左右上下四条边界
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	num, target := 1, n*n
	for num <= target {
		for i := l; i <= r; i++ { // left to right
			matrix[t][i] = num
			num += 1
		}
		t += 1
		for i := t; i <= b; i++ { // top to bottom
			matrix[i][r] = num
			num += 1
		}
		r -= 1
		for i := r; i >= l; i-- { // right to left
			matrix[b][i] = num
			num += 1
		}
		b -= 1
		for i := b; i >= t; i-- { // bottom to top
			matrix[i][l] = num
			num += 1
		}
		l += 1
	}
	return matrix
}
