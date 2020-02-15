package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func mapAbout() {
	// 映射将键映射到值。
	// 映射的零值为 nil 。nil 映射既没有键，也不能添加键。
	var m map[string]Vertex

	// make 函数会返回给定类型的映射，并将其初始化备用。
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat:  40.353,
		Long: -74.343,
	}
	fmt.Println(m["Bell Labs"])

	// 修改映射
	//m["Bell Labs"].Lat = 23.342  // 不行
	m["Bell Labs"] = Vertex{40.32, 44.32} //这个地方不能省略 Vertex
	fmt.Println(m["Bell Labs"])

	// 当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	fmt.Println(m["Google"])

	// 删除元素
	delete(m, "Bell Labs")
	fmt.Println(m["Bell Labs"])

	// 通过双赋值检测某个键是否存在, 不存在elem是该元素类型的零值 且 ok 为 false
	elem, ok := m["Bell Labs"]
	fmt.Println("value: ", elem, "Present: ", ok)

}

func mapLiterals() {
	/*
		var m = map[string]Vertex{
			"Bell Labs": Vertex{34.432, 43.43},
			"Google":    Vertex{32.223, 32.435},
		}
	*/

	// 文法中可以省略类型名
	var m = map[string]Vertex{
		"Bell Labs": {34.432, 43.43},
		"Google":    {32.223, 32.435},
	}
	fmt.Println(m)

}

func main() {
	mapAbout()
	mapLiterals()
}
