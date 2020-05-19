package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
10张票, 最终卖出了多于10张票
总结:
	不要以共享内存的方式去通信, 而已以通信的方式共享内存
*/

func main() {
	var ticket = 10
	go sellTicket1("窗口1", &ticket)
	go sellTicket1("窗口2", &ticket)
	go sellTicket1("窗口3", &ticket)
	go sellTicket1("窗口4", &ticket)
	time.Sleep(10 * time.Second)
	fmt.Println("main function run over")

}

func sellTicket1(window string, ticket *int) {
	rand.Seed(time.Now().UnixNano())
	var nums = 0
	for {
		if *ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(window, "sell ticket")
			*ticket--
			nums++
		} else {
			fmt.Println(window, "sold out")
			break
		}
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("%s 共售出%d张票\n", window, nums)
}
