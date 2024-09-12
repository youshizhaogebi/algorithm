package main

import (
	"fmt"
	"math/rand"
)

// 定义单链表结构体
type ListNode struct {
	Data int
	Next *ListNode
}

// 添加节点到链表末尾
func (head *ListNode) AddNode(data int) *ListNode {
	newNode := &ListNode{Data: data, Next: nil}
	if head == nil {
		return newNode
	} else {
		// 找到最后一个节点
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	return head
}

// 返回随机节点值
func (head *ListNode) GetRandom() int {
	index := 1
	selectPoint := head.Data // 初始化为第一个节点的值
	curr := head
	for curr != nil { // 遍历链表，selectPoint 会是随机节点中存储的值
		if rand.Float64() < 1.0/float64(index) {
			selectPoint = curr.Data
		}
		index++
		curr = curr.Next
	}
	return selectPoint
}

func main() {
	// 初始化链表
	var head *ListNode

	// 添加节点到链表
	head = head.AddNode(1)
	head = head.AddNode(2)
	head = head.AddNode(3)
	head = head.AddNode(4)
	head = head.AddNode(5)

	// 获取并打印随机节点值
	ret := head.GetRandom()
	fmt.Println(ret)
}
