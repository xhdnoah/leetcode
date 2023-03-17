(function (rowSum, colSum) {
    const row = rowSum.length, col = colSum.length
    const matrix = new Array(row).fill(0).map(() => new Array(col).fill(0))
    let i = 0, j = 0
    while (i < row && j < col) {
        const v = Math.min(rowSum[i], colSum[j])
        matrix[i][j] = v
        rowSum[i] -= v
        colSum[j] -= v
        if (rowSum[i] === 0) {
            ++i
        }
        if (colSum[j] === 0) {
            ++j
        }
    }
    return matrix
})()