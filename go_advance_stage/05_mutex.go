package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
10张票, 最终卖出了多于10张票
总结:
	不要以共享内存的方式去通信, 而已以通信的方式共享内存
 */

var mutex sync.Mutex  // 创建锁
var wg sync.WaitGroup
func main() {
	var ticket = 10
	wg.Add(4)

	go sellTicket("窗口1", &ticket)
	go sellTicket("窗口2", &ticket)
	go sellTicket("窗口3", &ticket)
	go sellTicket("窗口4", &ticket)

	wg.Wait()
	fmt.Println("main function run over")

}

func sellTicket(window string, ticket *int) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	var nums = 0
	for {
		// 上锁
		mutex.Lock()
		if *ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(window, "sell ticket")
			*ticket--
			nums++
		} else {
			mutex.Unlock()
			fmt.Println(window, "sold out")
			break
		}
		// 解锁
		mutex.Unlock()
	}
	time.Sleep(500*time.Millisecond)
	fmt.Printf("%s 共售出%d张票\n", window, nums)
}
