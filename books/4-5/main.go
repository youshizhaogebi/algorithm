package main

import (
	"fmt"
	"sort"
)

// 人由 (a, b) 组成， a 表示身高，b 表示该人前面身高大于或等于他的人数
// 思路：
// 最高的人排一队，第二高的根据 b 值重新排队，以此类推

// 接收 people 切片，
func reorderQueue(people [][]int) [][]int {
	// 使用 sort.Slice 函数对 people 进行排序。排序规则如下：
	// 先比较身高 a（即 people[a][0] 和 people[b][0]），身高较小的排在前面。
	// 如果两个身高相等，则比较 b（即前面有多少人比自己高），b 值较小的排在前面。
	sort.Slice(people, func(a, b int) bool { // 匿名函数 func() 返回 true，元素 a 排在 b 前面
		if people[a][0] == people[b][0] { // 比较身高，矮的在前
			return people[a][1] < people[b][1]
		} else {
			return people[a][0] < people[b][0]
		}
	})

	// 结果返回 ans 切片，初始化为{{-1, -1}, {-1, -1}, ...}
	ans := make([][]int, len(people))
	for i := range ans {
		ans[i] = []int{-1, -1}
	}

	// 根据b值插入
	// 外层循环遍历排序后的people。
	// index表示当前人在插入时所期望的前面身高大于或等于自己的人数。
	// 内层循环遍历ans，寻找合适的位置进行插入：
	// 如果当前位置未被占用且index为0，则将当前人插入该位置。
	// 如果当前位置未被占用或者当前位置的身高和当前人相等，index减1，继续寻找插入点。
	// 这个过程确保了较高的人在队列中插入时，前面有足够多的高个子。
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		for j := 0; j < len(people); j++ {
			if ans[j][0] == -1 && index == 0 {
				ans[j][0], ans[j][1] = people[i][0], people[i][1]
				break
			} else if ans[j][0] == -1 || ans[j][0] == people[i][0] {
				index--
			}
		}
	}

	// 返回结果切片
	return ans
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	ret := reorderQueue(people)
	fmt.Println(ret)
}
