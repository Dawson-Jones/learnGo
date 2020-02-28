package utils

import "fmt"

func Count()  {
	fmt.Println("Count function in utils package")
}

/*
init()函数: go的保留函数, 不能有参数和返回值, 在main函数执行前执行
一个包多个init函数, 排序从上至下执行
多个包的init函数, 按照导入顺序执行
import _ "package" 仅调用该包的init函数, 不使用其实的函数

 */
func init()  {
	fmt.Println("init function in utils package, for init some stuff")
}