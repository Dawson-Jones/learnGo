package main

import (
	"fmt"
	"reflect"
)

func structTagParse() {
	type T struct {
		A int    `json:"aaa" test:"testaaa"`
		B string `json:"bbb" test:"testbbb"`
	}

	t := T{
		A: 123,
		B: "hello",
	}
	tt := reflect.TypeOf(t)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Println(json)
		}
		test := field.Tag.Get("test")
		fmt.Println(test)
	}
}

func convertType() {
	type T struct {
		A int    `newT:"AA"`
		B string `newT:"BB"`
	}

	type newT struct {
		AA int
		BB string
	}

	t := T{
		A: 123,
		B: "hello",
	}
	tt := reflect.TypeOf(t)
	tv := reflect.ValueOf(t)

	var new_t newT
	new_t_v := reflect.ValueOf(&new_t)
	fmt.Println(new_t_v)

	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		newTTag := field.Tag.Get("newT")
		new_t_v.Elem().FieldByName(newTTag).Set(tv.Field(i))
	}

	fmt.Println(new_t)
}

func main() {
	// struct tag 解析
	//structTagParse()

	// 类型转换和赋值
	convertType()
}
