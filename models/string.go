package models

import()

// 字符串相关操作

// 反转字符串
func ReverseString(s []byte) {
	// 直接交换 s[i], s[len-i-1]
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

