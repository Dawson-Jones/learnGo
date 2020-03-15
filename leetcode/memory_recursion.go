package main

import "fmt"

// fibonacci series
func fib(n int) int {
	dp := make([]int, 1000)
	//dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	res := fibRecur(n, dp)
	fmt.Println(dp[:n+1])
	return res
}

func fibRecur(n int, dp []int) int {
	if dp[n] !=0{
		return dp[n]
	}
	pre1 := fibRecur(n-1, dp)
	pre2 := fibRecur(n-2, dp)
	res := pre1 + pre2
	dp[n] = res
	return dp[n]
}

func main() {
	res := fib(7)
	fmt.Println(res)
}
