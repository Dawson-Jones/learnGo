package main

import "fmt"

func main() {
	/*
	   // C 语言中
	   struct point {
	   	int x;
	   	int y;
	   }
	   struct point pt;
	   pt = {1, 2};
	*/
	type Vertex struct {
		X int
		Y int
	}

	// 初始结构体
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	// 结构体指针
	p := &v
	(*p).X = 100
	fmt.Println(*p)

	// 结构体文法
	var (
		v1    = Vertex{1, 2}
		v2    = Vertex{X: 1} // 隐式赋予 Y: 0
		v3    = Vertex{}     // X, Y 都是 0
		point = &Vertex{1, 2}
	)
	fmt.Println(v1, point, v2, v3)
}
