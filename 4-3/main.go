package main

import (
	"fmt"
)

// 判断一组不同值的 push 和 pop 是否是一个空栈一系列操作后的结果
// 将 push数组 依次入栈，找出 pop数组 栈顶元素后出栈
// 最终遍历完 pop数组，说明清空整个栈

// 根据给定的 push数组、 pop数组 清空栈
func validateStack(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	// 循环将 pushed数组 入栈，每入栈一个与 pop数组 比较，若是 pop数组 中的首个元素，则匹配上
	for _, x := range pushed {
		stack = append(stack, x)
		// 若和 pop数组 匹配上，就出栈一个，在同样的位置循环匹配
		for len(stack) != 0 && j < len(pushed) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1] // 出栈一个
			j++
		}
	}
	return j == len(pushed)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 5, 4, 3, 1}
	ret := validateStack(arr1, arr2)
	fmt.Println(ret)
}
