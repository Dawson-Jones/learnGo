package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x) // float64
	v := reflect.ValueOf(x)
	fmt.Println("type: ", t)
	fmt.Printf("value: %v, type: %T\n", v, v) // value: 3.4, type: reflect.Value  类型不是float64了
	// --------------------------------
	fmt.Println("kind is float64", v.Kind() == reflect.Float64) // true
	fmt.Println(v.Type())                                       // float64
	fmt.Println(v.Float())                                      // 3.4
}
