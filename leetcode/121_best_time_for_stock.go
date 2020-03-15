package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	curMaxBenefit := 0
	if len(prices) == 0{
		return curMaxBenefit
	}
	curMinPrice := prices[0]
	for _, v:= range prices{
		curMinPrice = int(math.Min(float64(curMinPrice), float64(v)))
		curMaxBenefit = int(math.Max(float64(curMaxBenefit), float64(v-curMinPrice)))
	}
	return curMaxBenefit
}

func main() {
	prices := []int{7,1,5,3,6,4}
	benefit := maxProfit(prices)
	fmt.Println(benefit)
}