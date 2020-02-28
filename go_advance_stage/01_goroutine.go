package main

import "fmt"

func printNum()  {
	for i:=0;i<100;i++{
		fmt.Printf("num: %d\n", i)
	}
}

func main() {
	go printNum()  // 主函数结束 协程会结束, 不管有没有运行完
	for i:=0;i<100;i++{
		fmt.Printf("char A %d\n", i)
	}

}
