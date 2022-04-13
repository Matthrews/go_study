// 当前程序包名
package main

// 导入其他包

// 调用的时候只需要Println()，而不需要fmt.Println()
import . "fmt"

// 起个别名

// 多个同时导入

// 常量声明
const PI = 3.14

// 全局变量声明与赋值
var name = "goopher"

// 类型声明
type NewType int

// 结构声明
type golang interface{}

// main函数为程序入口
func main() {
	Println("Hello World!")
}

// 函数名首字母小写即为private
func getId() {}

// 反之，首字母大写几位public
func GetName() {}
