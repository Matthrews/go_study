// 数据类型
// 布尔型，数字型，字符串型，派生类型（Pinter, Array, Struct, Channel, Func, Slice, Map, Interface）
package main

import "fmt"

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的变量声明

func test() {
	var available bool // 一般声明
	valid := true      // 简短声明
	available = true   // 赋值操作

	fmt.Println(available, valid)
}

func main() {
	test()
	fmt.Println(isActive, enabled, disabled)
}
