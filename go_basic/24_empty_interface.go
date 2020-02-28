package main

import "fmt"

/*
指定了零个方法的接口值被称为 *空接口：*

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
 */

type E interface {
}

type Cat struct {
	color string
}

type Person struct {
	name string
	age int
}

func test1(e E)  {
	fmt.Println(e)
}

func test2(e interface{})  {
	fmt.Println("--->: ", e)
}

func main() {
	//var c Cat = Cat{"白猫"}
	c := Cat{"白猫"}
	var e1 E = Cat{color:"花猫"}
	var e2 E  = Person{
		name: "Dawson",
		age:  20,
	}
	var e3 E = "I'm OK"
	var e4 E = 121

	fmt.Println(c.color)
	//fmt.Println(e1)
	//fmt.Println(e2)
	//fmt.Println(e3)
	//fmt.Println(e4)
	test1(e1)
	test1(e2)
	test1("call me lover boy")
	test1(23)

	test2(e3)
	test2(e4)
}
