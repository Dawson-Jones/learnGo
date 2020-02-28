/*
非缓冲通道: make(chan 类型)
	一次发送, 一次接收, 都会阻塞
缓冲通道: make(chan 类型, 缓冲容量)
	发送: 缓冲区满了才阻塞
	接收: 缓冲区空了才阻塞
 */
package main

import (
	"fmt"
	"strconv"
)

func seeDifferent()  {
	ch1Nobuf := make(chan int)
	ch2Buf := make(chan int, 5)

	fmt.Println(len(ch1Nobuf), cap(ch1Nobuf)) // 0,0
	fmt.Println(len(ch2Buf), cap(ch2Buf))     // 0,5
}
func aSimpleEg()  {
	ch := make(chan int, 2)
	ch<-1
	ch<-2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
func sendData(ch chan string)  {
	for i:=0;i<10;i++{
		ch<-"数据"+strconv.Itoa(i)
		fmt.Printf("goroutine write %d datum\n", i)
	}
	close(ch)
}

func example()  {
	ch := make(chan string, 5)
	go sendData(ch)
	for v:= range ch{
		fmt.Println("\tI read data:", v)
	}
	fmt.Println("that's all")
}

func main() {
	//seeDifferent()
	//aSimpleEg()
	//example()

}

