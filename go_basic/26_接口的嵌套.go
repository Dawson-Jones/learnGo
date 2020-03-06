// 类似于类中的继承

package main

import "fmt"

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
	name string
	age int
	color string
}

func (c Cat) test1()  {
	fmt.Printf("the cat named %s\n", c.name)
}
func (c Cat) test2()  {
	fmt.Printf("the cat is %d year old\n", c.age)
}
func (c Cat) test3()  {
	fmt.Printf("woo, a %s cat\n", c.color)
}

func main() {
	var cat = Cat{
		name:  "芝麻糊",
		age:   2,
		color: "gray",
	}
	cat.test1()
	cat.test2()
	cat.test3()
	// -------------
	var a A = cat
	a.test1()  // A 接口的实例只能实现A接口的方法
	var b B = cat
	b.test2()
	var c C  = cat
	c.test1()
	c.test2()
	c.test3()
}