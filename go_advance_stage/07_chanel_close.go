package main

import "fmt"

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	for v := range ch {
		fmt.Println("get data:", v)
	}
	fmt.Println("main function is over")

}
