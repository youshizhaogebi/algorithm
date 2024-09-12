package main

import "fmt"

// 切片反转
func reverse(slice []int) {
	// 循环交换切片第一个和最后一个的值
	for x, y := 0, len(slice)-1; x < y; x, y = x+1, y-1 {
		slice[x], slice[y] = slice[y], slice[x] // 在同一行会互换，写在两行会按顺序执行
	}
}

func main() {
	slice := []int{1, 6, 168}
	reverse(slice)
	fmt.Println(slice)
}
