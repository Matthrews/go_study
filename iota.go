// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

package main

import (
	"fmt"
	"unsafe"
)

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// 等效写法
const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a, b, c)

	fmt.Println("<<==========================>>")

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // iota +=1
		g        // 100  iota +=1
		h = iota // 7,恢复计数
		i        // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

	fmt.Println("<<==========================>>")

	const (
		o = 1 << iota
		p = 3 << iota
		m
		n
	)

	fmt.Println("o =", o)
	fmt.Println("p =", p)
	fmt.Println("m =", m)
	fmt.Println("n =", n)

	fmt.Println("<<==========================>>")

	var x = "hello"
	fmt.Println("Sizeof: ", unsafe.Sizeof(x))

	fmt.Println("<<==========================>>")

	// iota 只是在同一个 const 常量组内递增，每当有新的 const 关键字时，iota 计数会重新开始
	const (
		s = iota
		u = iota
		v = iota
	)
	const xx = iota
	const yy = iota
	fmt.Println(s, u, v, xx, yy)
	// 输出是 0 1 2 0 0
}
