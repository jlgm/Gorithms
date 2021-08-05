package main

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	markX, markY := make([]bool, n), make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				markX[i] = true
				markY[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if markX[i] || markY[j] {
				matrix[i][j] = 0
			}
		}
	}
}
