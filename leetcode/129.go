package main

/*
给定一个二叉树，它的每个结点都存放一个0-9的数字，每条从根到叶子节点的路径都代表一个数字。
例如，从根到叶子节点路径 1->2->3 代表数字 123。
计算从根到叶子节点生成的所有数字之和。
说明:叶子节点是指没有子节点的节点。

输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.
*/
func Travel(vertex *TreeNode, curSum int, res *int) {
	if vertex == nil {
		return
	}
	if vertex.Left == nil && vertex.Right == nil {
		*res += curSum*10 + vertex.Val
	}
	Travel(vertex.Left, curSum*10+vertex.Val, res)
	Travel(vertex.Right, curSum*10+vertex.Val, res)
}

func sumNumbers(root *TreeNode) int {
	var res = 0
	Travel(root, 0, &res)
	return res
}

func main() {

}
