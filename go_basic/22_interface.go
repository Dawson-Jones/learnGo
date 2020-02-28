package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type myFloat float64  // type 定义的类型是一个全新的类型, 而C中typedef定义的是已经存在的类型

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := myFloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	// Vertex指针类型实现了接口方法, 值类型没有实现, 所以不能使用下面的情况
	// a = v
	// fmt.Println(a.Abs())
}

