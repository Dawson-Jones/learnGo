/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
 */

package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		var curRes []string
		for j := 0; j < i; j++ {
			for _, preV := range dp[j] {
				for _, aftV := range dp[i-1-j] {
					curCase := "(" + preV + ")" + aftV
					curRes = append(curRes, curCase)
				}
			}
		}
		dp[i] = curRes
	}
	return dp[n]
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
