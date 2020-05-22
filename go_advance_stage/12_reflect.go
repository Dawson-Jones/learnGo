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
	rfValue := reflect.ValueOf(p1)
	fmt.Printf("value: %v, type: %T\n", rfValue, rfValue)
	method:= rfValue.MethodByName("PrintInfo")
	fmt.Printf("method: %v, type: %T\n", method, method)
	fmt.Println("methodValue-kind:", method.Kind(), "-----type:", method.Type())
	// 无参数调用
	method.Call(nil)
	// 或者
	args := make([]reflect.Value, 0)
	method.Call(args)
	// -------------
	method= rfValue.MethodByName("Say")
	fmt.Println(method.Kind(), method.Type()) // func func(string)
	args0 := []reflect.Value{reflect.ValueOf("world"), reflect.ValueOf(1)}
	method.Call(args0)
}
