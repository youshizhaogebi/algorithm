package models

import()

// 哈希表相关操作

// 返回两个数组的交集
func Intersection(nums1 []int, nums2 []int) []int {
	count := make(map[int]bool)
	for _, num := range nums1 {
		count[num] = true
	}
	result := []int{}
	for _, num := range nums2 {
		if count[num] {
			result = append(result, num)
			delete(count, num) // 删除已有元素，避免重复
		}
	}
	return result
}

// 找出和为目标值的两个数，返回下标
func TwoSum(nums []int, target int) []int {
	// 哈希表，一层循环，另一个数在表中寻找
	// 降低时间复杂度
	indexMap := make(map[int]int)
	for i, num := range nums {
		complementNum := target - num
		if idx, found := indexMap[complementNum]; found {
			return []int{idx, i}
		}
		// 哈希表中记录当前数字和下标
		indexMap[num] = i
	}
	return nil
}
