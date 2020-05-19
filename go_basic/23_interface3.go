package main

import "fmt"

type I interface {
	M()
}

type T struct {
	s string
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

func main() {
	var i I
	var t *T
	i = t
	i.M()
	describe(i)

	i = &T{"hello"}
	i.M()
	describe(i)
}
