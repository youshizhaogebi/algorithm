package models

import (
	"container/list"
)

// 二叉树相关操作

// 链式存储的二叉树
// 比链表多一个指针

// 定义二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreOrderTraversal(root *TreeNode) []int {
	// 方法一：迭代法，切片实现栈
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root} // 初始化栈，根节点入栈
	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问该节点
		result = append(result, node.Val)

		// 右子树入栈，左子树入栈。左子树会先处理
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result

	// 方法二：递归法

	// var result []int
	// // 定义递归函数
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	// 根节点开始
	// 	result = append(result, node.Val)
	// 	// 遍历左子树
	// 	traverse(node.Left)
	// 	// 遍历右子树
	// 	traverse(node.Right)
	// }
	// traverse(root)
	// return result
}

// 中序遍历
func InOrderTraversal(root *TreeNode) []int {
	// 方法一：迭代法，使用双向链表作为栈
	result := []int{}
	stack := list.New() // 使用双向链表作为栈
	current := root
	for current != nil || stack.Len() > 0 {
		// 先遍历左子树
		for current != nil {
			stack.PushBack(current)
			current = current.Left
		}

		// 访问栈顶节点
		if stack.Len() > 0 {
			node := stack.Remove(stack.Back()).(*TreeNode) // 弹出栈顶元素
			result = append(result, node.Val)              // 访问节点
			current = node.Right                           // 继续遍历右子树
		}
	}
	return result

	// // 方法二：递归法

	// var result []int
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	traverse(node.Left)
	// 	result = append(result, node.Val)
	// 	traverse(node.Right)
	// }
	// traverse(root)
	// return result
}

// 定义操作类型，用于区分“访问”还是“处理”
type Command struct {
	Action string // "visit" or "process"
	Node   *TreeNode
}

// 后序遍历
func PostOrderTraversal(root *TreeNode) []int {
	// 方法一：统一迭代法
	if root == nil {
		return nil
	}
	result := []int{}
	stack := []Command{{"visit", root}}
	for len(stack) > 0 {
		// 弹出栈顶元素
		command := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if command.Action == "process" {
			// 如果是处理操作，记录节点值
			result = append(result, command.Node.Val)
		} else if command.Node != nil {
			// 后序遍历的顺序是 左 -> 右 -> 根
			stack = append(stack, Command{"process", command.Node}) // 处理当前节点
			if command.Node.Right != nil {
				stack = append(stack, Command{"visit", command.Node.Right}) // 访问右子树
			}
			if command.Node.Left != nil {
				stack = append(stack, Command{"visit", command.Node.Left}) // 访问左子树
			}
		}
	}
	return result

	// 方法二：递归法

	// var result []int
	// var traverse func(node *TreeNode)
	// traverse = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	traverse(node.Left)
	// 	traverse(node.Right)
	// 	result = append(result, node.Val)
	// }
	// traverse(root)
	// return result
}

// 层次遍历，广度优先搜素（一层一层，从左到右）
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	curLevel := []*TreeNode{root}
	for len(curLevel) > 0 {
		// 预分配每一层节点的值
		vals := make([]int, 0, len(curLevel))
		nextLevel := make([]*TreeNode, 0, len(curLevel)*2) // 下一次节点数最多是当前层两倍

		// 遍历当前层节点
		for _, node := range curLevel {
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		// 添加当前层的节点值到结果中
		res = append(res, vals)
		// res = append([][]int{vals}, res...) // 自底向上层次遍历。还可以遍历 res 交换位置

		curLevel = nextLevel
	}

	return res
}

// Morris 遍历。修改树的结果，空间复杂度 O(1)
func MorrisInOrder(root *TreeNode) []int {
	result := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
			current = current.Right
		} else {
			// 找到当前节点左子树的最右节点
			pre := current.Left
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
			}

			// 建立临时链接
			if pre.Right == nil {
				pre.Right = current
				current = current.Left
			} else {
				// 访问当前节点
				pre.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

// 判断两个树是否相同
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归

	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// 判断 root 中是否包含 subRoot
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 递归
	
	if root == nil {
		return false
	}
	if IsSameTree(root, subRoot) {
		return true
	}
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
