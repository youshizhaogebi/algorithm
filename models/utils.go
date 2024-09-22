package models

import()

// 工具类函数

// 返回绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 返回两数中最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回两数中最小值
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}