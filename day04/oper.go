// 1. go的特殊运算符

// & 返回变量存储地址
// * 指针变量

// 2. 运算符优先级

// 由左至右由高到低如下：

// * / % << >> &
// + - | ^
// == != < <= > >=
// &&
// ||

// 3. 指针变量 * 和地址值 & 的区别
package main

import "fmt"

func main()  {
	var a int = 4
	var ptr *int
	ptr = &a

	fmt.Println("a的值为：", a)
	fmt.Println("*ptr为：", *ptr)
	fmt.Println("ptr为：", ptr)

	// 4. Go 的自增，自减只能作为表达式使用，而不能用于赋值语句。
	a++
	a--
	b = a++  // 编译报错

	// 5. 使用指针变量与不使用的区别：
	var a int = 4
    var ptr int
    ptr = a 
    fmt.Println(ptr)//4
    a = 15
    fmt.Println(ptr)//4
    
    var b int = 5 
    var ptr1 *int
    ptr1 = &b 
    fmt.Println(*ptr1)//5
    b=15 
    fmt.Println(*ptr1)//15

}
