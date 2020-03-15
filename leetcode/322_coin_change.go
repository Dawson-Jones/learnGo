/*
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
*/

package main

import (
	"fmt"
	"math"
)

// 纯递归
func coinChange(coins []int, amount int) int {
	var dp []int
	dp = append(dp, 0)
	for i := 1; i <= amount; i++ {
		pieces := 1 << 30
		for _, v := range coins {
			if i-v >= 0 {
				pieces = int(math.Min(float64(pieces), float64(dp[i-v])+1))
			}
		}
		dp = append(dp, pieces)
	}
	for i, v := range dp {
		if v >= 1<<30 {
			dp[i] = -1
		}
	}
	return dp[amount]
}

func main() {
	typeCoin := []int{2, 4}
	res := coinChange(typeCoin, 5)
	fmt.Println(res)
}
