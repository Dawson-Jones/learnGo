package main

import "math"

func minEnergy(matrix [][]int) int {
	length := len(matrix)

	res := make([][]int, length)
	for i := range res {
		res[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			res[i][j] = math.MaxInt32
		}
	}
	res[0][0] = 0

	for i := 1; i < length; i++ {
		res[0][i] = res[0][i-1] + int(math.Abs(float64(matrix[0][i-1]-matrix[0][i])))
		res[i][0] = res[i-1][0] + int(math.Abs(float64(matrix[i-1][0]-matrix[i][0])))
	}

	for i := 1; i < length; i++ {
		for j := i; j < length; j++ {
			res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i][j-1])+math.Abs(float64(matrix[i][j-1]-matrix[i][j]))))
			res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i-1][j])+math.Abs(float64(matrix[i-1][j]-matrix[i][j]))))
		}
	}

	return res[length-1][length-1]
}

func main() {
	matrix := [][]int{
		{1, 4, 4, 4, 4},
		{1, 4, 1, 1, 1},
		{1, 4, 1, 4, 1},
		{1, 4, 1, 4, 1},
		{1, 1, 1, 4, 1},
	}

	println(minEnergy(matrix))
}
