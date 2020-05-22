package main

import "fmt"

type Slice []int

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Println(elem)
	return s
}

func main() {
	// defer 推迟 推迟的函数调用会被压入一个栈中。
	// 当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}

	s := make(Slice, 0)
	defer s.Add(1).Add(2) // defer 只会推迟最后一个函数的执行, 这里Add(1) 没有被延迟执行
	s.Add(3)
	fmt.Println(s)
	fmt.Println("done")
}
