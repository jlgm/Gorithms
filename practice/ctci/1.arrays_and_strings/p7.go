package main

func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
	return matrix
}

// func main() {
// 	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	fmt.Println(rotate(a))
// }
