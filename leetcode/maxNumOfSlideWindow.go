package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	length := len(nums)
	if length==0{
		return res
	}
	st := 0
	ed := k
	for ed<=length{
		var currentMax = 1<<31*-1
		for i:=st;i<ed;i++{
			if nums[i]>currentMax{
				currentMax = nums[i]
			}
		}
		res = append(res, currentMax)
		st++
		ed++
	}
	return res
}

func main() {
	var nums = []int{1,3,-1,-3,5,3,6,7}
	res := maxSlidingWindow(nums, 3)
	fmt.Println(res)
}
