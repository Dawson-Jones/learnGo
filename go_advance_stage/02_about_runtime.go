package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	// 获取逻辑cpu数量
	fmt.Println("nums of logic CPU:", runtime.NumCPU())
	// 设置go程序使用最大的cpu数量 [1, 256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)
}

func goroutine1() {
	defer fmt.Println("goroutine1 defer...")
	//runtime.Goexit()
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine1:", i)
	}
}

func goroutine2() {
	for i := 0; i < 5; i++ {
		// 让出时间片, 让其他goroutine先执行
		runtime.Gosched()
		fmt.Println("goroutine2:", i)
	}
}

func main() {
	// 获取goroot目录
	fmt.Println("GOROOT -->", runtime.GOROOT())
	// 获取操作系统
	fmt.Println("current OS: ", runtime.GOOS)

	// gosched
	//go goroutine1()
	//go goroutine2()

	// goexit  终止当前的goroutine
	go func() {
		fmt.Println("goroutine start")
		goroutine1()
		fmt.Println("goroutine end") // 没有执行, 因为goroutine1终止了当前的goroutine
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("main function run over")

}
