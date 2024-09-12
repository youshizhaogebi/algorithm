package main

import (
	"fmt"
	"reflect"
)

type Json struct {
	tag string `json:"Tag"`
}

func printTag(input interface{}) {
	t := reflect.TypeOf(input).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体字段 %v 对应的 json tag 是%v\n", t.Field(i).Name, t.Field(i).Tag.Get("json"))
	}
}

func main() {
	j:=Json{
		tag:"test",
	}
	printTag(&j)
}
