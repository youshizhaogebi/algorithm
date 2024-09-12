package qa

import ()

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// 假设无重复数字

// 示例 1:
// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4 （9在 nums 中且下标为4，从0开始数）

func BinarySearch(nums []int, target int) int {
	// 初始化左右边界
	left := 0
	right := len(nums) - 1

	// 循环，缩小范围
	for left <= right { // target 范围是 [) 时用 <
		// 求中点
		mid := left + (right-left)>>1 // right - left 不会产生溢出，右移 1 位相当于除以 2

		// 判断，调整区间
		if nums[mid] == target { // 找到了
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1 // target 范围 [) 时用 right = middle
		}
	}

	// 找不到
	return -1
}
