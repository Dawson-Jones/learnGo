package main

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// 函数
	fmt.Println(add(2, 3))

	// 多值返回
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 命名返回值
	fmt.Println(split(17))
}
