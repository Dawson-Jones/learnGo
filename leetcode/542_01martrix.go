/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。

输入:
0 0 0
0 1 0
1 1 1

输出:
0 0 0
0 1 0
1 2 1
*/
package main

import "fmt"

func getOneLevel(r, c, rowNums, colNums int) [][]int {
	top := r - 1
	btm := r + 1
	lft := c - 1
	rht := c + 1
	var levelArr [][]int
	if top >= 0 {
		levelArr = append(levelArr, []int{top, c})
	}
	if btm < rowNums {
		levelArr = append(levelArr, []int{btm, c})
	}
	if lft >= 0 {
		levelArr = append(levelArr, []int{r, lft})
	}
	if rht < colNums {
		levelArr = append(levelArr, []int{r, rht})
	}
	return levelArr
}

func updateMatrix(matrix [][]int) [][]int {
	rowNums := len(matrix)
	if rowNums == 0 {
		return [][]int{{}}
	}
	colNums := len(matrix[0])
	resMartrix := make([][]int, rowNums)

	for r, rowArr := range matrix {
		rowArrRes := make([]int, colNums)
		for c, value := range rowArr {
			level := 0
			if value != 0 {
				var queue [][][]int
				oneLevel := getOneLevel(r, c, rowNums, colNums)
				queue = append(queue, oneLevel)

				for len(queue) != 0 {
					level++
					firstArr := queue[0]
					queue = queue[1:]
					var totalArr [][]int
					for _, arr := range firstArr {
						if matrix[arr[0]][arr[1]] != 0 {
							oneLevel = getOneLevel(arr[0], arr[1], rowNums, colNums)
							for _, v:= range oneLevel{
								totalArr = append(totalArr, v)
							}
						} else {
							totalArr = [][]int{}
							break
						}
					}
					if len(totalArr) > 0 {
						queue = append(queue, totalArr)
					}
				}
			}
			rowArrRes[c] = level
		}
		resMartrix[r] = rowArrRes
	}
	return resMartrix
}

func main() {
	martrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	res := updateMatrix(martrix)
	fmt.Print(res)
}
