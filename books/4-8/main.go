package main

import "fmt"

// 设计最近最少使用 (LRU) 缓存约束数据结构
// LRUCache(int capacity): 初始化 LRU 缓存
// int get(int key): 判断 key 是否存在
// void put(int key, int value): 添加键值对到缓存中
// get() put() 时间复杂度 O(1)
// 思路：
// 双向链表

// LRUCache 结构体用于表示最近最少使用 (LRU) 缓存
// 它包含一个容量限制，一个双向链表的头和尾节点，以及一个哈希表用于存储键到节点的映射
type LRUCache struct {
	capacity   int
	head, tail *Node
	values     map[int]*Node
}


// 节点
type Node struct {
	key, value int
	prev, next *Node
}

// 用于初始化 LRUCache
// capacity 参数指定缓存的最大容量
func Constructor(capacity int) LRUCache {
	return LRUCache{
		values:   map[int]*Node{},
		capacity: capacity,
	}
}

// 从缓存中获取指定键的值
// 如果键存在，则将对应节点移动到链表的尾部，并返回该值
// 如果键不存在，则返回 -1
func (lr *LRUCache) Get(key int) int {
	node, ok := lr.values[key]
	if !ok {
		return -1
	}
	lr.moveToLast(node)
	return node.value
}

// 将指定节点移动到链表的尾部
func (lr *LRUCache) moveToLast(node *Node) {
	if node == lr.tail {
		return
	}
	if node == lr.head {
		lr.head = lr.head.next
		lr.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	lr.tail.next = node
	node.prev = lr.tail
	lr.tail = lr.tail.next
	lr.tail.next = nil
}

// 将一个新的节点添加到链表的尾部
func (lr *LRUCache) append(key, value int) {
	node := &Node{
		key:   key,
		value: value,
	}
	if lr.tail == nil {
		lr.tail = node
		lr.head = node
	} else {
		lr.tail.next = node
		node.prev = lr.tail
		lr.tail = node
	}
	lr.values[key] = node
}

// 将一个键值对放入缓存中
// 如果键已经存在，则更新值，并将节点移动到链表的尾部
// 如果键不存在且缓存未满，则将新节点添加到链表的尾部
// 如果缓存已满，则删除链表的头节点（最少使用的节点），并将新节点添加到链表的尾部
func (lr *LRUCache) Put(key int, value int) {
	if _, ok := lr.values[key]; ok {
		lr.values[key].value = value
		lr.moveToLast(lr.values[key])
		return
	}
	if len(lr.values) < lr.capacity {
		lr.append(key, value)
		return
	}
	node := lr.head
	lr.moveToLast(node)
	node.key = key
	node.value = value
}

func main() {
	obj := Constructor(2)
	obj.Put(5, 88)
	res := obj.Get(5)
	fmt.Println(res)
}
