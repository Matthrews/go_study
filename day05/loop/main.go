package main

import "fmt"

func main()  {
	// 1. numbers
	i := 0
	for i < 5 {
		fmt.Println("NumberOne", i)
		i++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("NumberTwo", i)
	}

	// 2. strings

	names := []string{ "yoshi", "mario", "peach", "bowser", "luigi" }
	
	for i := 0; i < len(names); i++ {
		fmt.Println("NameOne", names[i])
	}

	for index, name := range names {
		fmt.Println("NameTwo", index, name)
	}

	for _, name := range names {
		// doesn't alter orignal string value
		name = "new string"
		fmt.Println("NameThree", name)
	}
	fmt.Println("NameThree", names)
	


}