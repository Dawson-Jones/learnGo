package main

import (
	"fmt"
	"math"
)

/*
二叉树的最大宽度

输入:

          1
         /
        3
       / \
      5   3

输出: 2
解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root== nil {
		return 0
	}
	curLevelPos := []int{1}
	var ntxLevelPos []int
	var curLevelVertex = []*TreeNode{root}
	var nxtLevelVertex []*TreeNode
	maxWidth := 1
	length := 1

	for length > 0 {
		length = 0
		for i, v := range curLevelVertex {
			if v.Left != nil {
				ntxLevelPos = append(ntxLevelPos, curLevelPos[i]*2)
				nxtLevelVertex = append(nxtLevelVertex, v.Left)
				length++
			}
			if v.Right != nil {
				ntxLevelPos = append(ntxLevelPos, curLevelPos[i]*2+1)
				nxtLevelVertex = append(nxtLevelVertex, v.Right)
				length++
			}
		}
		if length > 0{
			if length > 1{
				maxWidth = int(math.Max(float64(maxWidth), float64(ntxLevelPos[length-1]-ntxLevelPos[0]) + 1))
			}
			curLevelPos = ntxLevelPos
			ntxLevelPos = []int{}
			curLevelVertex = nxtLevelVertex
			nxtLevelVertex = []*TreeNode{}
		}
	}
	return maxWidth
}

func main() {
	//t15 := TreeNode{
	//	5,nil, nil,
	//}
	//t10 := TreeNode{
	//	5,nil, nil,
	//}
	//t9 := TreeNode{
	//	5,nil, nil,
	//}
	//t8 := TreeNode{
	//	5,nil, nil,
	//}
	//t7 := TreeNode{
	//	5,nil, &t15,
	//}
	//t6 := TreeNode{
	//	5,nil, nil,
	//}
	t5 := TreeNode{
		5,nil, nil,
	}
	t4 := TreeNode{
		4,nil, nil,
	}
	//t3 := TreeNode{
	//	3,nil, &t7,
	//}
	t2 := TreeNode{
		2, &t4, &t5,
	}
	t1 := TreeNode{
		1, &t2, nil,
	}
	res := widthOfBinaryTree(&t1)
	fmt.Print(res)
}
