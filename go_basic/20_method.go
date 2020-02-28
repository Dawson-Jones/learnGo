package main

import (
	"fmt"
	"math"
)

/*
Go 没有类。不过你可以为结构体类型定义方法。
方法就是一类带特殊的 接收者 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
*/
type Structure struct {
	X, Y float64
}

func (v Structure) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Structure) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v Structure) ScaleWithoutPoint(f float64) {
	v.X *= f
	v.Y *= f
}
func (v Structure) Compute() Structure {
	v.X *= 10
	v.Y *= 10
	return v
}

func AbsFunc(v Structure) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// 方法
	v := Structure{X: 3, Y: 4}
	fmt.Println(v.Abs())

	// 方法指针
	// 使用指针接收者的原因有二：
	//
	// 1. 方法能够修改其接收者指向的值。
	//
	// 2. 这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

	// 如果不用指针 v的值并没有改变 因为参数的传递是把值复制了一份, 所以堆栈函数并没有改变原调用方的值
	v.ScaleWithoutPoint(10)
	fmt.Println(v.Abs())
	// 使用指针, v的值改变
	v.Scale(10)
	fmt.Println(v.Abs())

	v = Structure{3, 4}
	fmt.Println(v.Compute())
	fmt.Println(AbsFunc(v)) // 5

	p := &Structure{3, 4}

	// 因为Compute接收的是一个Structure类型的值, 而不是指针 所以这里做了隐式的变换
	// 即: (*p)找到这个类型的值, 然后复制给了Compute函数, 所以 *p 的值没有变化
	fmt.Println(p.Compute())
	fmt.Println(AbsFunc(*p)) // 5

}
