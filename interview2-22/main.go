package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
goroutine 调用顺序
*/

func main() {
	runtime.GOMAXPROCS(1) // 最大处理器为 1

	// 等待 10 个 goroutines 执行
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 5; i++ {
		// i 被 goroutine 共享，i 的值不确定，可能重复
		go func() {
			fmt.Println("i: ", i)
			wg.Done() // 使  WaitGroup 内部的计数器减 1
		}() // 没有传入参数
	}
	for j := 0; j < 5; j++ {
		// 最后创建的 goroutine 会最先输出，打印 j: 4
		go func(i int) {
			fmt.Println("j: ", i)
			wg.Done()
		}(j) // 传入 j 的当前值
	}
	wg.Wait() // 阻塞，直到 WaitGroup 内部的计数器变为 0
}
