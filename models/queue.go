package models

import ()

// 队列相关操作
// 数组队列，还有

// 1. 普通切片队列

// 定义队列结构体
type Queue struct {
	data []int
}

// 入队
func (q *Queue) Push(value int) {
	q.data = append(q.data, value)
}

// 出队
func (q *Queue) Pop() int {
	front := q.data[0]
	q.data = q.data[1:]
	return front
}

// 返回队列第一个元素
func (q *Queue) Peek() int {
	return q.data[0]
}

// 2. 单调递减队列。第一个元素是最大值
// 用于
// 只维护可能成为最大值的元素。Pop() 和 Push() 操作时进行判断

// 定义单调递减队列结构体
type MonotonicQueue struct {
	data []int
}

// 创建单调递减队列
func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		data: make([]int, 0),
	}
}

// 入队
func (mq *MonotonicQueue) Push(value int) {
	// 移除队列中小于 value 的元素
	for len(mq.data) > 0 && mq.data[len(mq.data)-1] < value {
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, value)
}

// 出队
func (mq *MonotonicQueue) Pop(value int) {
	// 判断是否要移除队列中的最大值
	if len(mq.data) > 0 && mq.data[0] == value {
		mq.data = mq.data[1:]
	}
}

// 返回队列中最大值（首个元素）
func (mq *MonotonicQueue) Max() int {
	return mq.data[0]
}

// 返回在数组中移动的窗口中的最大值
func MaxSlidingWindow(nums []int, k int) []int {
	// 单调队列。入队、出队后，保持队列首位是最大值
	if len(nums) == 0 || k <= 0 {
		return nil
	}

	mq := NewMonotonicQueue()
	res := make([]int, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 入队当前元素
		mq.Push(nums[i])

		// 从第 k 个元素开始记录最大值
		if i >= k-1 {
			// 将窗口的最大值存入结果
			res[i-k+1] = mq.Max()
			// 移除窗口左边界的元素
			mq.Pop(nums[i-k+1])
		}
	}

	return res
}

// 3. 优先级队列

