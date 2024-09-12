package main

/*
接口循环调用

Programmer 结构体有一个 String() 方法，该方法尝试打印 Programmer 实例。
当调用 p.String() 时，fmt.Sprintf 内部尝试将 p 转换为字符串表示，这通常会调用 String() 方法（如果类型定义了此方法）。
导致 String() 方法被递归调用，因为没有明确的终止条件，最终会因为栈溢出而崩溃。
要解决这个问题，应该避免在 String() 方法中直接调用 fmt.Sprintf 来格式化包含自身的结构体。
你可以直接返回你想要的字符串格式。
*/

// type Programmer struct {
// 	Name string
// }

// func (p *Programmer) String() string {
// 	return fmt.Sprintf("print: %v", p)
// }

// func main() {
// 	p := &Programmer{}
// 	p.String()
// }
