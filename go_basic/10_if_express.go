package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" // 数字和字符串相加默认转换为字符串
	}
	// return math.Sqrt(x) // 因为这个是float64类型所以不能返回
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if 语句可以在条件表达式前执行一个简单的语句。该语句声明的变量作用域仅在 if else之内。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))
	fmt.Println(pow2(3, 2, 10), pow2(3, 3, 10))
}
