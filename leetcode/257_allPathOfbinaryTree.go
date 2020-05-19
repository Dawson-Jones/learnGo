package main

import "strconv"

func travel(vertex *TreeNode, res *[]string, curString string) {
	if vertex==nil {
		return
	}
	if vertex.Left == nil && vertex.Right == nil {
		*res = append(*res, curString + "->"+ strconv.Itoa(vertex.Val))
		return
	}
	travel(vertex.Left, res, curString+"->"+strconv.Itoa(vertex.Val))
	travel(vertex.Right, res, curString+"->"+strconv.Itoa(vertex.Val))
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		return append(res, strconv.Itoa(root.Val))
	}
	travel(root.Left, &res, strconv.Itoa(root.Val))
	travel(root.Right, &res, strconv.Itoa(root.Val))
	return res
}

func main() {

}
