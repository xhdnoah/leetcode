package main

// board =
// [
//   ['A','B','C','E'],
//   ['S','F','C','S'],
//   ['A','D','E','E']
// ]
// Given word = "SEE", return true.
// Given word = "ABCB", return false.
// 在地图上的任意一个起点开始，向 4 个方向分别 DFS 搜索
// 直到所有的单词字母都找到了就输出 true，否则输出 false
var direction = [][]int{
	{-1, 0}, // 向上
	{0, 1},  // 向右
	{1, 0},  // 向下
	{0, -1}, // 向左
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) // 边界探测
}

// index 表示 DFS 搜索到第几步
func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + direction[i][0]
			ny := y + direction[i][1]
			// 注意 !visited[nx][ny] 为防止搜索回到原点
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		// 一轮 DFS 结束，置为默认
		visited[x][y] = false
	}
	return false
}
