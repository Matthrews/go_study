package main

import "fmt"

func main() {
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	for index, name := range names {
		if index == 1 {
			fmt.Println("continue at ", index)
			continue
		}

		if index > 2 {
			fmt.Println("break at ", index)
			break
		}

		fmt.Printf("The value at %v is %v\n", index, name)
	}
}
