package main

import (
	"fmt"
	"time"
)

func useTimer(){
	timer:= time.NewTimer(3*time.Second)  // 创建一个计时器, 3秒后触发 写入通道  Timer类型
	fmt.Println(time.Now())
	ch:=timer.C
	fmt.Println(<-ch)
}
func timerStop()  {
	timer := time.NewTimer(5*time.Second)
	go func() {
		<-timer.C
		fmt.Println("run other function")
	}()
	time.Sleep(3*time.Second)
	signal := timer.Stop()
	if signal{
		fmt.Println("stop timer successful")
	}
}
func timerAfter()  {
	ch:=time.After(3*time.Second)
	fmt.Println(time.Now())
	<-ch
	fmt.Println(time.Now())

}
func main() {
	//useTimer()
	// timer stop
	//timerStop()
	// timer after 直接返回就是通道, 就是newTimer.C
	timerAfter()
}
