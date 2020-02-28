package main

import (
	"fmt"
	"math"
)

func main() {
	// 类型判断
	v := 42
	fmt.Printf("v is of type %T\n", v)

	// 类型转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	// var xx string = string(f)  // cannot convert f (type float64) to type string
	fmt.Println(x, y, f, z)
}
