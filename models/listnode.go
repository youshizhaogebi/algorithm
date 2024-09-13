package models

// 操作单链表

// 单链表：
// 构造链表时创建虚拟头节点
// 添加节点时创建新节点
// 查询时创建新指针

// 定义节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 存储链表信息
type MyLinkedList struct {
	dummyHead *ListNode // 虚拟头节点
	Size      int       // 链表大小
}

// 初始化链表
func Constructor() MyLinkedList {
	newNode := &ListNode{
		0,
		nil,
	}
	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

// 查询
func (list *MyLinkedList) Get(index int) int {
	if list == nil || index < 0 || index >= list.Size {
		return -1
	}
	// cur 是第一个真正的头节点
	cur := list.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

// 添加头节点
func (list *MyLinkedList) AddAtHead(val int) {
	// 创建新节点
	newNode := &ListNode{Val: val}
	newNode.Next = list.dummyHead.Next
	list.dummyHead.Next = newNode
	list.Size++
}

// 添加最后一个节点
func (list *MyLinkedList) AddAtTail(val int) {
	// 创建新节点
	newNode := &ListNode{Val: val}

	// 查询，创建新指针
	cur := list.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
	list.Size++
}

// 插入特定位置
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		index = 0
	} else if index > list.Size {
		return
	}

	// 插入，创建新节点
	// 遍历，创建新指针
	newNode := &ListNode{Val: val}
	cur := list.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode
	list.Size++
}

// 删除特定节点
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= list.Size {
		return
	}

	cur := list.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	list.Size--
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	// cur 遍历过程中，利用 tmp、pre 实现交换
	// cur 移动后，将当前节点指向 pre
	cur := head
	var pre *ListNode = nil // pre 初始化为 nil，不会打印出来，不然会打印 0
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// 两两交换节点
func SwapPairs(head *ListNode) *ListNode {
	// 使用虚拟头节点，画图弄清楚如何交换
	dummyHead := &ListNode{
		Next: head,
	}
	pre := dummyHead
	for head != nil && head.Next != nil {
		// 要交换的节点S
		firstNode := head
		secondNode := head.Next

		// 交换
		// 第一个指针指向第二个的后一个
		// 第一个的前一个指向第二个
		// 第二个指向第一个
		pre.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		// 移动节点
		pre = firstNode
		head = firstNode.Next
	}
	return dummyHead.Next
}

// 删除链表倒数第 n 个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 快慢指针，慢指针距离 n，快指针到尾部时删除慢指针指向的节点
	dummyHead := &ListNode{
		Next: head,
	}
	slow, fast := dummyHead, dummyHead
	for i := 0; i <= n; i++ { // 慢指针需要指向第 n-1 个节点
		fast = fast.Next
	}
	for fast != nil { // n+1 次循环后 fast 可能指向 nil，无法访问 fast.Next
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
