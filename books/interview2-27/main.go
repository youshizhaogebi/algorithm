package main

import (
	"fmt"
)

// 只会打印 closed，ch 通道在 goroutine 启动前就关闭了

func main() {
	ch := make(chan int, 100)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	go func() {
		for i := range ch {
			fmt.Println("i: ", i)
		}
	}()
	close(ch)
	fmt.Println("closed")
}
