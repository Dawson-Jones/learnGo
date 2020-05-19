package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	currentLevel := 0
	totalLevel := 0
	var levelVertex [][]*TreeNode
	firstVertex := []*TreeNode{root}
	levelVertex = append(levelVertex, firstVertex)
	totalLevel++

	for currentLevel != totalLevel {
		var tempVertexes []*TreeNode
		var curLevelValues []int
		for _, curVertex := range levelVertex[currentLevel] {
			curLevelValues = append(curLevelValues, curVertex.Val)
			if curVertex.Left != nil {
				tempVertexes = append(tempVertexes, curVertex.Left)
			}
			if curVertex.Right != nil {
				tempVertexes = append(tempVertexes, curVertex.Right)
			}
		}
		res = append(res, curLevelValues)
		currentLevel++
		if tempVertexes != nil {
			levelVertex = append(levelVertex, tempVertexes)
			totalLevel++
		}
	}
	for k := 1; k < len(res); k += 2 {
		arr := res[k]
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return res
}

func main() {
	vertex4 := TreeNode{4, nil, nil}
	vertex3 := TreeNode{3, nil, &vertex4}
	vertex2 := TreeNode{2, nil, nil}
	vertex1 := TreeNode{1, &vertex2, &vertex3}
	res := levelOrderBottom(&vertex1)
	for _, vs := range res {
		fmt.Println(vs)
	}
	//fmt.Println(res)
}
