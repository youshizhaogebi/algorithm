package main

import "fmt"

/*
FrequencyStack 频率栈
push(int x) 压入栈；
pop() 删除出现频次最高的元素（若频率相同，删除靠近栈顶的元素）
*/

// 频次栈
type FrequencyStack struct {
	freq    map[int]int   // [x]频次 保存相同频次数
	group   map[int][]int // [频次][x, x, ...] 保存相同频次组，用于频次相同的情况
	maxfreq int           
}

// 新建栈，创建两个空 map
func NewFrequencyStack() FrequencyStack {
	hash := make(map[int]int)
	maxHash := make(map[int][]int)
	return FrequencyStack{freq: hash, group: maxHash}
}

// 入栈，更新频次，更新频次组
func (fs *FrequencyStack) Push(x int) {
	// 入栈值 x 的频次加 1
	fs.freq[x]++

	// 若出现频次更高，修改最大频次值，并将 x 加到对应的频次组里
	f := fs.freq[x]
	if f > fs.maxfreq {
		fs.maxfreq = f
	}
	fs.group[f] = append(fs.group[f], x) // [x, x, ...]
}

// 出栈，减少频次，更新频次组
func (fs *FrequencyStack) Pop() int {
	tmp := fs.group[fs.maxfreq] // [x, x, ...]
	x := tmp[len(tmp)-1] // 最大频次组中，切片的最后一个 x
	fs.group[fs.maxfreq] = fs.group[fs.maxfreq][:len(fs.group[fs.maxfreq])-1] // 最大频次组弹出最后一个数
	fs.freq[x]-- // x 的频次减 1
	if len(fs.group[fs.maxfreq]) == 0 {
		fs.maxfreq--
	}
	return x
}

func main() {
	i := 8
	j := 16
	obj := NewFrequencyStack()
	obj.Push(i)
	obj.Push(j)
	ret1 := obj.Pop()
	ret2 := obj.Pop()
	fmt.Println(ret1)
	fmt.Println(ret2)
}
