package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {

}

func main() {
	vertex4 := TreeNode{4, nil, nil}
	vertex3 := TreeNode{3, nil, &vertex4}
	vertex2 := TreeNode{2, nil, nil}
	vertex1 := TreeNode{1, &vertex2, &vertex3}
	res := zigzagLevelOrder(&vertex1)
	for _, vs := range res {
		fmt.Println(vs)
	}
	//fmt.Println(res)
}
