package main

import(
	"fmt"
)

// 求数组元素和

// 返回数组元素和
func sumArray(array [5]int) int{
	sum:=0
	for _, val := range array {
		sum += val
	}
	return sum
}

func main(){
	array:=[5]int{1,2,3,4,5}
	total:=sumArray(array)
	fmt.Println(total)
}