package main

import (
	"fmt"
	"math"
	"strings"
)

func sayHello(str string) {
	fmt.Printf("Good morning, %v\n", str)
}

func sayBye(str string) {
	fmt.Printf("Goodbye, %v\n", str)
}

func patch(strs []string, callback func(string)) {
	for _, str := range strs {
		callback(str)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, name := range names {
		initials = append(initials, name[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	sayHello("mario")
	sayHello("luigi")
	sayBye("luigi")

	patch([]string{"cloud", "tifa", "bareet"}, sayHello)

	s := circleArea(10.5)
	fmt.Println("CircleArea", s)
	fmt.Printf("CircleArea is %0.3f\n", s)

	fn, ln := getInitials("tifa lockhart")
	fmt.Println("getInitials", fn, ln)

	fnTwo, lnTwo := getInitials("coffee hat")
	fmt.Println("getInitials", fnTwo, lnTwo)

	fnThree, lnThree := getInitials("banana")
	fmt.Println("getInitials", fnThree, lnThree)
}
