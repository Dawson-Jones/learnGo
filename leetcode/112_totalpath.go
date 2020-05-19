package main

func travel(vertex *TreeNode, sum int, cursum int) bool {
	if vertex == nil {
		return false
	}
	if vertex.Left == nil && vertex.Right == nil{
		if sum == cursum + vertex.Val{
			return true
		} else {
			return false
		}
	}
	if travel(vertex.Left, sum, cursum + vertex.Val){
		return true
	}
	if travel(vertex.Right, sum, cursum + vertex.Val){
		return true
	}
	return false
}

func hasPathSum(root *TreeNode, sum int) bool {
	return travel(root, sum, 0)
}
