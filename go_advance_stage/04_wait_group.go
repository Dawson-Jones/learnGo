package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func fun(s string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("func %s\n", s)
	}
}
func main() {
	wg.Add(2)
	go fun("A")
	go fun("B")

	fmt.Println("main function is blocking")
	wg.Wait()
	fmt.Println("main function has been unblocked ")

}
