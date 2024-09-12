package qa

import (
	"testing"
)

// 用例1
func TestBinarySearch(t *testing.T) {

	// 定义测试用例结构体
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 3, 5, 7, 9}, 7, 3},
	}

	// _ 忽略索引，遍历 testCases 切片，
	for _, tc := range testCases {
		got := BinarySearch(tc.arr, tc.target)
		if got != tc.want {
			t.Errorf("search(%v, %d) = %d; want %d", tc.arr, tc.target, got, tc.want)
		}
	}
}
