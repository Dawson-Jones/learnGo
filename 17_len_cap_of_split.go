package main

import (
	"fmt"
)

func splitShow() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// length 0, capacity 6
	s = s[:0]
	printSlice(s)
	// length 4, capacity 6
	s = s[:4]
	printSlice(s)
	// length 2, capacity 4
	s = s[2:]
	printSlice(s)
	// 拓展了长度
	s = s[:4]
	printSlice(s)

	// 切片的零值是 nil, nil 切片的长度和容量为 0 且没有底层数组。
	s = s[4:]
	fmt.Println("s: ")
	printSlice(s)
	nilOrNo(s)
	var s2 []int
	fmt.Println("s2: ")
	printSlice(s)
	nilOrNo(s2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func nilOrNo(s []int) {
	if s == nil {
		fmt.Printf("%v: nil\n", s)
	} else {
		fmt.Printf("%v: not nil\n", s)
	}
}

func makeSplit()  {
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	// make([]类型, 长度, 容量)
	a := make([]int, 5)
	fmt.Println("a: ")
	printSlice(a)

	b := make([]int, 0, 5)
	fmt.Println("b: ")
	printSlice(b)

	c := b[:2]
	fmt.Println("c: ")
	printSlice(c)

	d := c[2:5]
	fmt.Println("d: ")
	printSlice(d)

}
func main() {
	// 切片
	splitShow()

	// 用 make 切片
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	makeSplit()

}
