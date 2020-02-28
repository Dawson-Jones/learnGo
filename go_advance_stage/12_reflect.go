package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string, digit int) {
	fmt.Println("hello,", msg, digit)
}

func (p Person) PrintInfo() {
	fmt.Printf("name: %s, age: %d, sex: %s\n", p.Name, p.Age, p.Sex)
}
func main() {
	p1 := Person{"Dawson", 20, "male"}
	fmt.Printf("p1: %v, type: %T\n", p1, p1)
	value := reflect.ValueOf(p1)
	fmt.Printf("value: %v, type: %T\n", value, value)
	methodValue := value.MethodByName("PrintInfo")
	fmt.Printf("methodValue: %v, type: %T\n", methodValue, methodValue)
	fmt.Println("methodValue-kind:", methodValue.Kind(), "-type:", methodValue.Type())
	// 无参数调用
	methodValue.Call(nil)
	// 或者
	args := make([]reflect.Value, 0)
	methodValue.Call(args)
	// -------------
	methodValue = value.MethodByName("Say")
	fmt.Println(methodValue.Kind(), methodValue.Type()) // func func(string)
	args0 := []reflect.Value{reflect.ValueOf("world"), reflect.ValueOf(1)}
	methodValue.Call(args0)

}
