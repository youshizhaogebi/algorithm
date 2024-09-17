package models

import "strconv"

// 栈相关操作

// 需要先实例化后，再进行操作，例如：主函数中 stack := MyStack{} ，再进行 stack.Push(x)
// Size(), Empty() 操作方便链表实现的栈。因此有必要实现这两个操作

// 1. 切片实现栈

// 切片实现栈
type MyStack []int

// 入栈，添加数据到末尾
func (s *MyStack) Push(v int) {
	*s = append(*s, v)
}

// 弹出栈，返回弹出的数值
func (s *MyStack) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

// 查看栈顶数值
func (s *MyStack) Peek() int {
	return (*s)[len(*s)-1]
}

// 返回栈长度
func (s *MyStack) Size() int {
	return len(*s)
}

// 判断栈是否为空
func (s *MyStack) Empty() bool {
	return len(*s) == 0
}

// 返回通过逆波兰表达式计算的数值
func EvalRPN(tokens []string) int {
	// 数字入栈，表达式则取出栈顶两个数计算后再入栈
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token) // 字符串转为整数
		if err == nil {
			// token 为整数，入栈
			stack = append(stack, val)
		} else {
			// token 不为整数，如果为符号，取栈顶两个数计算，结果入栈
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

// 2. 链表实现栈
