package main

import "fmt"

func getNext(n int) int {
	var res int
	for n > 0{
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func isHappy(n int) bool {
	hashMap := make(map[int]int)
	for n != 1{
		if hashMap[n] == 1{
			return false
		}
		hashMap[n] = 1
		n = getNext(n)
	}
	return true
}

func main() {
	res := isHappy(19)
	fmt.Println(res)
}
