package main

import "fmt"


func trap(height []int) int {
	var stack  = []int{0}  // 存放的是 height[] 的 index
	stackLastIndex := 0
	for i, v:= range height[1:] {
		if stackLastIndex >= 0 &&
	}

}

func main() {
	req := []int{
		0,1,0,2,1,0,1,3,2,1,2,1,
	}
	res := trap(req)
	fmt.Println(res)
}
