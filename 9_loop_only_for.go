package main

import "fmt"

func main() {
	for1()
	for2()
	for_is_while()
	
}

func for1()  {
	// var sum int = 0 
	var sum = 0 
	for i := 0; i <10; i++{
		sum += i
	}
	fmt.Println(sum)
	
}

func for2()  {
	sum := 1
	for ; sum < 1000;{
		sum += sum
	}
	fmt.Println(sum)
	
}

func for_is_while()  {
	sum := 1
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)
}

func loop_forever()  {
	// 三个条件什么都不写就是死循环
	for {
	}
}