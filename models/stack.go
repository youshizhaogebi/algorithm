package models

import "strconv"

// 栈相关操作

// 需要先实例化后，再进行操作，例如：主函数中 stack := Stack{} ，再进行 stack.Push(x)
// Size(), Empty() 操作方便链表实现的栈。因此有必要实现这两个操作

// 切片实现栈
type Stack struct {
	data []int
}

// 入栈，添加数据到末尾
func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

// 弹出栈，返回弹出的数值
func (s *Stack) Pop() int {
	val := (s.data)[len(s.data)-1]
	s.data = (s.data)[:len(s.data)-1]
	return val
}

// 查看栈顶数值
func (s *Stack) Peek() int {
	return (s.data)[len(s.data)-1]
}

// 返回栈长度
func (s *Stack) Size() int {
	return len(s.data)
}

// 判断栈是否为空
func (s *Stack) Empty() bool {
	return len(s.data) == 0
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
