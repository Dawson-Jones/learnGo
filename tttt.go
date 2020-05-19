package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Unix(time.Now().Unix(), 0)
	fmt.Println(t)
}
