package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chanTest1()
	time.Sleep(10 * time.Second)
}

func chanTest1() {
	defer fmt.Println("父协程退出")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 10; i++ {
		go ctxTest2(ctx, i)
	}
	time.Sleep(5 * time.Second)
}
