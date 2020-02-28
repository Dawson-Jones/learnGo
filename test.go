package main

import (
	"fmt"
	"strings"
)

// 单词出现次数
func wordCount(s string) map[string]int {
	m := make(map[string]int)
	arr := strings.Split(s, " ")
	for _, value := range arr {
		if _, ok := m[value]; ok {
			m[value] += 1
		} else {
			m[value] = 1
		}
	}
	return m
}

// 斐波那契数列, 闭包
func fibonacci() func() int {
	pre := 0
	after := 1
	return func() int {
		current := pre
		pre = after
		after += current
		return current
	}
}



func main() {
	m := wordCount("call me lover boy")
	fmt.Println(m)

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}





