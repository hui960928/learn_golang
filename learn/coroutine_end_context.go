package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxTest1()
	time.Sleep(10 * time.Second)
}

func ctxTest1() {
	defer fmt.Println("父协程退出")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 10; i++ {
		go ctxTest2(ctx, i)
	}
	time.Sleep(5 * time.Second)
}

func ctxTest2(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("子协程接受到停止信号", num)
			return
		default:
			fmt.Println("子协程执行中", num)
			fmt.Println("ChanTest2", num)
			time.Sleep(1 * time.Second)
		}
	}
}
