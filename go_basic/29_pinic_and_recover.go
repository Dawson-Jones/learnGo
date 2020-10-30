/*
### panic:
	1. 内建函数
	2. 加入函数 F 中书写的 panic 语句, 会终止其后要执行的代码,
	   panic 所在函数 F 内如果存在要执行的 defer 函数列表则逆序执行
	3. 返回函数的调用者 G, 在 G 中调用函数 F 语句之后的代码不会执行, G 中之前 defer 的函数列表也逆序执行
       defer 类似于 try-catch-finally 中的 finally

// 错误和异常是 err 和 panic

### recover:
	1. 内建函数
	2. 用来控制一个 goroutine 的 panicking 行为, 捕获 panic
	3. 调用建议:
	   - 在 defer 函数中, 通过 recover 终于一个 goroutine 的 panicking 过程, 从而恢复正常代码的执行
	   - 可以获取通过 panic 传递的 err
*/

package main

import "fmt"

func funA() {
	fmt.Println("i am function A")
}

func funB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "--------the program is recovered")
		}
	}()
	fmt.Println("i am function B")
	defer fmt.Println("first defer function B")
	for i := 1; i < 10; i++ {
		fmt.Println("i: ", i)
		if i == 5 {
			// 发生恐慌后, 已经被 defer 的函数还会执行, 接下来的就不执行了
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
