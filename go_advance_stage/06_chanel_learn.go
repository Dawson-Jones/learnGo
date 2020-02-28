/*
声明通道:
	var 通道名 chan 数据类型
创建通道: 如果通道为nil(就是不存在), 就需要先创建通道
	通道名 = make(chan 数据类型)
简短创建: c := make(chan, int)
*/

package main

import "fmt"

func aboutChanel() {
	var c chan int                                     // c: nil
	fmt.Printf("chanel c Type: %T\tValue: %v\n", c, c) // chanel c Type: chan int	Value: <nil>
	c = make(chan int)
	fmt.Printf("chanel c Type: %T\tValue: %v\n", c, c) // chanel c Type: chan int	Value: 0xc0000180c0

}

func communicateWithChanel1() {
	ch1 := make(chan bool)
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Printf("the %d time print\n", i)
		}
		ch1 <- true // 往Chanel中写数据
	}()

	signal := <-ch1 // 读取数据, 未读到之前是阻塞的
	fmt.Println("main signal", signal)
	fmt.Println("main function over")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func communicateWithChanel2() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("x: %d, y: %d, x + y: %d\n", x, y, x+y)
}
func main() {
	//aboutChanel()

	// 使用chanel通信
	//communicateWithChanel1()
	communicateWithChanel2()

}
