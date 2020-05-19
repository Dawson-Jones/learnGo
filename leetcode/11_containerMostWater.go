package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	afterIndex := len(height) - 1
	currentIndex := 0
	maxContain := 0
	for currentIndex <= afterIndex {
		curMax := int(math.Min(float64(height[currentIndex]), float64(height[afterIndex]))) * (afterIndex - currentIndex)
		maxContain = int(math.Max(float64(curMax), float64(maxContain)))
		if height[currentIndex] < height[afterIndex] {
			currentIndex++
		} else {
			afterIndex--
		}
	}
	return maxContain
}

func main() {
	req := []int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	}
	res := maxArea(req)
	fmt.Println(res)
}
