/*
### panic:
	1. 内建函数
	2. 加入函数F中书写的panic语句, 会终止其后要执行的代码,
	   panic所在函数F内如果存在要执行的defer函数列表则逆序执行
	3. 返回函数的调用者G, 在G中调用函数F语句之后的代码不会执行, G中之前defer的函数列表也逆序执行
       defer类似于 try-catch-finally 中的finally

// 错误和异常是err 和 panic

### recover:
	1. 内建函数
	2. 用来控制一个goroutine的panicking行为, 捕获panic
	3. 调用建议:
	   - 在defer函数中, 通过recover终于一个goroutine的panicking过程, 从而恢复正常代码的执行
	   - 可以获取通过panic传递的err
 */


package main

import "fmt"

func funA(){
	fmt.Println("i am function A")
}

func funB(){
	defer func() {
		if msg:=recover();msg != nil{
			fmt.Println(msg, "the program is recovered")
		}
	}()
	fmt.Println("i am function B")
	defer fmt.Println("first defer function B")
	for i:=1;i<10;i++{
		fmt.Println("i: ", i)
		if i==5{
			// 发生恐慌后, 已经被defer的函数还会执行, 接下来的就不执行了
			panic("function B panic")
		}
	}
	defer fmt.Println("second defer function B")
}

func main() {
	funA()
	defer fmt.Println("first defer function main")
	funB()
	defer fmt.Println("second defer function main")
}
