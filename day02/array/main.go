package main

import "fmt"

func main()  {

	// array
	var agesOne [3]int = [3]int{ 20, 34, 56 }
	var agesTwo= [3]int{ 20, 34, 56 }

	agesThree := [3]int{ 20, 34, 56 }

	fmt.Println("Println", agesOne)
	fmt.Println("Println", agesTwo)
	fmt.Println("Println", agesThree)

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println("Println", scores, len(scores))

	// slice ranges
	rangOne := scores[1:3]
	rangTwo := scores[1:]
	rangThree := scores[:3]
	fmt.Println("Println", rangOne, rangTwo, rangThree)

}