package main

import (
	"fmt"
)

// 合并 K 个已排序的链表
// 思路：
// 分治，先合并两个，再合并 K 个

// 定义单链表节点
type ListNode struct {
	Data int
	Next *ListNode
}

// 辅助函数：传入数组，构建链表
func buildLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	// 创建链表头节点
	head := &ListNode{Data: values[0]}
	current := head
	// 依次添加节点
	for _, value := range values[1:] {
		newNode := &ListNode{Data: value}
		current.Next = newNode
		current = newNode
	}
	return head
}

// 合并 l1，l2 两个链表节点，返回 l2 节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Data < l2.Data {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// 合并 K 个链表
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeTwoLists(left, right)
}

func main() {
	nodeList1 := buildLinkedList([]int{1, 4, 5})
	nodeList2 := buildLinkedList([]int{1, 3, 4})
	nodeList3 := buildLinkedList([]int{2, 6})
	ListsK := []*ListNode{nodeList1, nodeList2, nodeList3}
	ret := mergeKLists(ListsK)

	// 打印结果
	for ret.Next != nil {
		fmt.Println(ret.Next.Data)
		ret = ret.Next
	}
}
