package main

import "fmt"

func typeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string) // ok is ture
	fmt.Println(s, ok)
	f, ok := i.(float64) // ok is false
	fmt.Println(f, ok)   // f的类型是float64 的零值
	//f = i.(float64)      // 断言错误, 触发恐慌 panic
	//fmt.Println(f)
}

func typeJudgment(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("Twice %v is %v.\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long.\n", v, len(v)) // %q 是引用, 会自动加上引号
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

func main() {
	typeAssertion()

	typeJudgment("hello")
	typeJudgment(21)
	typeJudgment(true)
}
