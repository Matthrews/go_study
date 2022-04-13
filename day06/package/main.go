package main

import "fmt"

var score = 99.6 // 放到此处没问题

func main() {
	sayHello("mario")

	for _, v := range points {
		fmt.Println("Point", v)
	}

	// var score = 59.6 // 放到此处就是undefined

	showScore()
}

// go run day06/package/main.go day06/package/greetings.go
