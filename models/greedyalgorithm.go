package models

import "sort"

// 分发饼干。传入两个数组，返回满足人数
func FindContentChildren(g []int, s []int) int {
	// 小饼干先分给需求小的
	// 先遍历饼干，再遍历胃口
	
	sort.Ints(g) // 胃口从小到大排序
	sort.Ints(s)
	index := 0 // 记录满足的数量
	for i := 0; i < len(s); i++ {
		if index < len(g) && g[index] <= s[i] {
			index++
		}

		// 全部满足则退出循环
		if index == len(g) {
			break
		}
	}

	return index

	// // 大饼干优先满足需求大的
	// // 先遍历胃口

	// sort.Ints(g)
	// sort.Ints(s)
	// index := len(s) - 1 // 饼干数量
	// result := 0         // 满足人数
	// for i := len(g) - 1; i >= 0; i-- {
	// 	if index >= 0 && s[index] >= g[i] {
	// 		result++
	// 		index--
	// 	}
	// 	if result == len(s) || index < 0 {
	// 		break
	// 	}
	// }

	// return result
}

// 返回最长摆动子序列长度，子序列可以不连续，从原序列删除元素
func WiggleMaxLength(nums []int) int {
	// 方法一：贪心。序列排布后，找序列局部峰数（忽略山坡上的数）
	// 需要考虑不同坡度，是否忽略山坡（单调出现平坡，非单调出现平坡）
	// 时间复杂度 O(n)

	n := len(nums)
	if n < 2 {
		return n
	}

	result := 1
	prevDiff := 0 // 记录前一对数的差值。初始化为 0，处理 n<2 的情况
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) { // 等于 0 时是平坡的情况
			result++
			prevDiff = diff // 摆动变化时更新
		}
	}

	return result

	// 方法二：动态规划。树维护区间最大值
	// 时间复杂度 O(n^2) 或 O(nlogn) （使用树）
}

// 返回子序列最大和
func MaxSubArray(nums []int) int {
	// 贪心法 Kadane’s Algorithm
	// 遍历数组来找到以每个元素为结尾的最大子数组，并逐步更新全局的最大子数组和
	// 如果加上当前元素会减少当前的和，重新开始计算，即放弃之前的子数组和，直接从当前元素开始作为新的子数组的起点

	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
