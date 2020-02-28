package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	v:= reflect.ValueOf(x)
	fmt.Printf("value: %v, type: %T\n", v, v)  // value: 3.4, type: reflect.Value  类型不是float64了
	// --------------------------------
	fmt.Println("kind is float64", v.Kind()==reflect.Float64)
	fmt.Println(v.Type())  // float64
	fmt.Println(v.Float())  // 3.4 又变成了float64类型
}
