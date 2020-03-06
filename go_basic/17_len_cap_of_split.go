package main

import (
	"fmt"
	"strings"
)

func splitShow() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	// length 0, capacity 6
	s = s[:0]
	printSlice(s)
	// length 4, capacity 6
	s = s[:4]  // [2 3 5 7]
	printSlice(s)
	// length 2, capacity 4
	s = s[2:]  // [5 7]
	printSlice(s)
	// 拓展了长度
	s = s[:4]  // [5 7 11 13]
	printSlice(s)

	// 切片的零值是 nil, nil 切片的长度和容量为 0 且没有底层数组。
	s = s[4:]
	fmt.Println("s: ")
	printSlice(s)
	nilOrNo(s)  // no
	var s2 []int
	fmt.Println("s2: ")
	printSlice(s)
	nilOrNo(s2)  // yes
}

func makeSplit() {
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
	// make([]类型, 长度, 容量)
	a := make([]int, 5)
	fmt.Println("a: ")
	printSlice(a)  // len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	fmt.Println("b: ")
	printSlice(b)  // len=0 cap=5 []

	c := b[:2]
	fmt.Println("c: ")
	printSlice(c)

	d := c[2:5]
	fmt.Println("d: ")
	printSlice(d)

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

func sliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_",},
		[]string{"_", "_", "_",},
		[]string{"_", "_", "_",},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "O"
	board[1][0] = "X"
	board[0][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendMethod() {
	var temp []int
	temp = append(temp, 0)
	printSlice(temp)
	temp = append(temp, 1, 2, 3, 4)
	printSlice(temp)
}

func forRange() {
	temp := []int{1, 2, 4, 8, 16, 32, 64, 128,}

	// 可以将下标或值赋予 _ 来忽略它。
	// for i, _ := range pow
	// for _, value := range pow
	// 若你只需要索引，忽略第二个变量即可
	//for i := range pow
	for i, v := range temp {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func main() {
	// 切片
	//splitShow()

	// 用 make 切片
	// make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	//makeSplit()

	// 切片的切片
	//sliceOfSlice()

	// append 方法
	//appendMethod()

	// 对数组使用for 循环
	forRange()
}
