package models

import ()

// 字符串相关操作

// 反转字符串
func Reverse(s []byte) {
	// 直接交换 s[i], s[len-i-1]
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

// 反转字符串中的单词，单词前后可能存在多个空格
func ReverseWords(s string) string {
	// 方法一：调用库函数 strings.Fields 分隔单词，加入新的字符串
	// 方法二：原地反转，空间复杂度 O(1)
	// 0. 移除多余空格
	// 1. 反转字符串
	// 2. 反转每个单词
	str := []byte(s)
	n := len(s)
	slow := 0 // 慢指针

	// 移除空格
	for i := 0; i < n; i++ {
		if str[i] != ' ' {
			if slow != 0 {
				str[slow] = ' '
				slow++
			}
			for i < n && str[i] != ' ' {
				str[slow] = str[i]
				slow++
				i++
			}
		}
	}
	str = str[:slow]

	// 反转字符串
	Reverse(str)

	// 反转每个单词
	end := 0 // 单词末尾
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			Reverse(str[end:i])
			end = i + 1
		}
	}
	return string(str)
}

// KMP 算法：一个串中查找是否出现另一个串，返回第一个匹配的位置（从 0 开始）
func KMP(haystack string, needle string) int {
	// 不匹配时，记录已匹配的部分，避免从头匹配
	// 使用前缀表（next 数组），记录不匹配时重新匹配的位置
	// 前缀表第 2 位：前 i 个字符组成字符串的相同前后缀的最大长度
	// （前缀指不包含最后一个字符的所有以第一个字符开头的连续子串；后缀是指不包含第一个字符的所有以最后一个字符结尾的连续子串）
	// next 数组可以等于前缀表，可以是前缀表统一 -1；这里等于前缀表，跳转到 next[j-1]
	// 时间复杂度 O(n+m)

	// KMP 主函数和 buildNext 函数思想类似，前缀匹配，不匹配则退回 next[j-1]

	if len(needle) == 0 {
		return 0
	}

	next := buildNext(needle)
	j := 0 // needle 的下标
	for i := 0; i < len(haystack); i++ {
		// 不匹配时跳转
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		// 匹配时
		if haystack[i] == needle[j] {
			j++
		}

		// 找到匹配的子串
		if j == len(needle) {
			return i - j + 1
		}
	}

	return -1
}

// 构建 next 数组，求最长相同前后缀
func buildNext(needle string) []int {
	// 前缀匹配，不匹配则退回 next[j-1]
	next := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// 判断字符串是否由子串重复多次构成
func RepeatedSubstringPattern(s string) bool {
	// 方法一：KMP 算法：
	n := len(s)
	next := buildNext(s)         // next[i] 表示以 s[i] 为结尾的前缀的最长相同前后缀的长度
	l := next[n-1]               // l 是 s 的最长前缀后缀的长度，next[n-1] 表示整个字符串的最长前缀后缀的长度
	return l > 0 && n%(n-l) == 0 // len(s)-next(len(s)-1) 即为最小重复单元的长度
	// 最小重复单元：
	// if n%n-l == 0 {
	// 	return s[:n-l]
	// }

	// 方法二：拼接查找。拼接 s，去掉首尾，看是否仍包含 s。有额外空间开销
	// if len(s) == 0 {
	// 	return false
	// }
	// t := s + s
	// return strings.Contains(t[1:len(t)-1], s)
}
