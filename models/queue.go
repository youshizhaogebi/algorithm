package models

import "container/heap"

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

// 3. 优先级队列（堆）
// 定义堆结构
type Pair struct {
	Num  int
	Freq int
}

// 最小堆
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Freq < h[j].Freq } // 按照频率从小到大排序
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 返回数组中出现频率前 k 高的元素
func TopKFrequent(nums []int, k int) []int {
	// 方法一：优先级队列
	// 根据频率构建 map
	// 构建小堆顶，放入堆（完全二叉树）中，遇到频率高的就弹出堆顶
	// 倒叙构建 res 数组T

	// 1. 统计每个元素的出现频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 使用小根堆来维护前 k 个频率最大的元素
	h := &MinHeap{}
	heap.Init(h)
	for num, freq := range freqMap {
		heap.Push(h, Pair{Num: num, Freq: freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 3. 提取堆中的元素
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Pair).Num
	}

	return result

	// 方法二：构建 map 后调用 sort 包进行排序
	// res := []int{}
	// freqMap := map[int]int{}
	// for _, num := range nums {
	// 	freqMap[num]++
	// }
	// for key, _ := range freqMap {
	// 	res = append(res, key)
	// }
	// sort.Slice(res, func(i, j int) bool {
	// 	return freqMap[res[i]] > freqMap[res[j]]
	// })
	// return res[:k]
}
