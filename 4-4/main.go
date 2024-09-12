package main

import (
	"fmt"
)

// 栈实现队列：push(x), pop(), peek(), empty()

type Queue struct {
	Stack *[]int
	Queue *[]int
}

// 初始化队列
func NewQueue() Queue {
	tmp1, tmp2 := []int{}, []int{}
	return Queue{Stack: &tmp1, Queue: &tmp2}
}

// 将元素 x 推到队列最后
func (queue *Queue) Push(x int) {
	*queue.Stack = append(*queue.Stack, x)
}   

// 出队，从队列前面移除元素，返回该元素
func (queue *Queue) Pop() int {
	// 队列为空，栈转移到队列
	if len(*queue.Queue) == 0 {
		queue.convertQueue(queue.Stack, queue.Queue)
	}

	popped := (*queue.Queue)[len(*queue.Queue)-1]
	*queue.Queue = (*queue.Queue)[:len(*queue.Queue)-1]
	return popped
}

// 获取最前面的元素
func (queue *Queue) Peek() int {
	if len(*queue.Queue) == 0 {
		queue.convertQueue(queue.Stack, queue.Queue)
	}
	return (*queue.Queue)[len(*queue.Queue)-1]
}

// 栈元素弹出，压入队列
func (queue *Queue) convertQueue(s, q *[]int) {
	for len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		*q = append(*q, popped)
	}
}

// 判断队列是否为空
func (queue *Queue) Empty() bool {
	return len(*queue.Stack)+len(*queue.Queue) == 0
}

func main() {
	queue := NewQueue()
	ret1 := queue.Empty()
	x := 8
	y := 99
	queue.Push(x)
	queue.Push(y)
	ret2 := queue.Peek()
	ret3 := queue.Empty()
	fmt.Println(ret1)
	fmt.Println(ret2)
	fmt.Println(ret3)
}
