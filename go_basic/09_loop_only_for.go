package main

import "fmt"

func main() {
	for1()
	forIsWhile()
	loopForever()

}

func for1() {
	// var sum int = 0 
	var sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forIsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func loopForever() {
	// 三个条件什么都不写就是死循环
	for {
		fmt.Println(1)
	}
}
