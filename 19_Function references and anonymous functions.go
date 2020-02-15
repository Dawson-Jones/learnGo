// 函数引用 匿名函数 闭包

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64 ,float64)float64) float64 {
	/*
	fn func(float64 ,float64)float64  Golang
	double (*fn)(double, double)      C
	 */
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	// 类似于匿名函数
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	// 传递函数的引用
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// 闭包 -> 返回函数的引用
	f := adder()
	for i := 0;i<10;i++{
		fmt.Println(f(i))  // 做累加的时候, 会一直影响包外的 sum 值
	}
}
