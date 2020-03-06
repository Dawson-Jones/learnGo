package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {
	// 映射
	map1 := make(map[string]interface{})
	map1["name"] = "李小花"
	map1["age"] = 30
	ergou := Person{
		name: "李二狗",
		age:  38,
	}
	map1["friend"] = ergou
	fmt.Println("map1:", map1)

	// 数组
	array1 := make([]interface{}, 0, 10)
	array1 = append(array1, map1, "蛋蛋", ergou, 123)
	fmt.Println(array1)
	loopArray(array1)
}

func loopArray(array []interface{})  {
	for i, v := range array{
		fmt.Printf("value of %d is %v\n", i, v)
	}
}
