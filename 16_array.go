package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// 切片操作
	newPrimes := primes[1:4]
	fmt.Println(newPrimes)

	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	//char *names[4] = {"John", "Paul", "George", "Ringo"}; // 分配了四个 char 类型的指针

	fmt.Println(names)
	// 切片之后只是把指针给了x, y 修改的时候仍然会修改到底层的字符串
	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)
	y[0] = "XXX"
	fmt.Println(x, y)
	fmt.Println(names)

	// 没有长度的数组
	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
