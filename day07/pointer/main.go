package main

import "fmt"

func updateName(n string) {
	n = "wedge"
}

func updateNameTwo(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	updateName(name)
	fmt.Println("Memory address of name is ", &name)

	m := &name

	fmt.Println("Memory address of m is ", m)
	fmt.Println("Value at memory address is ", *m)

	fmt.Println("Name", name)

	// use pointer
	updateNameTwo(m)
	fmt.Println("Name", name)
}
