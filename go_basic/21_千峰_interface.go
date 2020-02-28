package main

import "fmt"

func main() {
	/*
		接口:
		接口是一组方法的签名
		当某个类型为这个接口中的所有方法提供了方法的实现, 就被成为实现接口
		接口和类型的实现关系, 是非侵入式

		1. 当需要接口类型的对象时, 可以使用任意现在类对象代替
		2. 接口对象不能访问实现类中的属性
		3. 不需要显示定义接口实现
			java中:
			class Mouse implements USB{}
	*/
	m := Mouse{name:"罗技"}
	f := FlashDisk{name:"三星"}
	testInterface(m)
	testInterface(f)

	m.start()
	m.end()
	fmt.Println(m.name)
	var usb USB = f
	usb.start()
	usb.end()
	f.deleteData()
	//usb.delete() 不可以
	//fmt.Println(usb.name)  不能做到, 因为接口对象不能访问实现类中的属性
}

// 定义接口
type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标点点点")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束工作")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "U盘进行数据存储")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "弹出...")
}

func (f FlashDisk) deleteData()  {
	fmt.Println(f.name, "在删除数据...")
}

// 测试函数
func testInterface(usb USB)  {
	usb.start()
	usb.end()

}
