package models

import "sort"

// 数组相关操作
// 常用双指针法

// 原地移除元素
func RemoveElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 二分查找
func BinarySearch(nums []int, target int) int {
	// 注意左右边界
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

	return -1
}

// 数组中找三个不同的元素，元素和为 0
func ThreeSum(nums []int) [][]int {
	// 方法一：哈希法保证不重复非常麻烦
	// 方法二：双指针法。先排序。右指针在最右。相加后，大了左移右指针，小了右移左指针
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 对另外两个数去重，指针移动后如果相同就去掉
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

// 数组中找到不重复四元组，和为 target
func FourSum(nums []int, target int) [][]int {
	// 双指针法，把 target-nums[a] 当作新的 target，求三元组
	// 多元组依此类推
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > target && (nums[i] > 0 || target > 0) {
			break
		}
		newTarget := target - nums[i]
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// 后面是求三元组
			if nums[j] > newTarget && (nums[j] > 0 || newTarget > 0) {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[j] + nums[left] + nums[right]
				if sum == newTarget {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right-1] == nums[right] {
						right--
					}
					left++
					right--
				} else if sum < newTarget {
					left++
				} else {
					right--
				}

			}
		}
	}
	return res
}