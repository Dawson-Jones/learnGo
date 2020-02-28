package main

import (
	"fmt"
	"runtime"
	"time"
)

func switch1st() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	// 每个 case 后面不需要跟break, 且 case 无需是常量
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switch2nd() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Printf("%T\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days later")
	default:
		fmt.Println("Too far away")
	}
}

func switchLikeIfElse() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	switch1st()
	switch2nd()
	switchLikeIfElse()
}
