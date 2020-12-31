package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//cc()
	//time.Sleep(1 * time.Second)
	//deadline()
	//timeout()
	value()
}

// 取消
func cc() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("cancel the goroutine")
					return
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println("master run over")
}

// deadline
func deadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// timeout
func timeout()  {
	ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// value
func value()  {
	f := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	ctx := context.WithValue(context.Background(), "language", "GO")

	f(ctx, "language")
	f(ctx, "color")
}