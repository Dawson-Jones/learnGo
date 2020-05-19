/*
给定一个二叉树，找出其最小深度。

给定二叉树 [3,9,20,null,null,15,7]，

     3
    / \
   9  20
 /   /  \
1  15    7
返回它的最小深度 3 。
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	var bfs = [][]*TreeNode{{root}}

	for _, level := range bfs {

	}

}

func main() {

}
