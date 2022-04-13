package main

import "fmt"

func updateName(n string) {
	n = "wedge"
}

func updateNameTwo(n string) string {
	n = "wedge"
	return n
}

func updateMenu(y map[string]float64) {
	y["pie"] = 7.68
}

// non-pointer values and pointer wrapper values

func main() {
	// group A types => strings, ints, booleans, floats, arrays, structs
	name := "tifa"

	updateName(name)
	fmt.Println("Name", name)

	name = updateNameTwo(name)
	fmt.Println("Name", name)

	// group B types => slices, maps, functions
	menu := map[string]float64{
		"soup":    9.45,
		"pie":     7.00,
		"vinegar": 8.54,
		"coffee":  32.00,
	}

	updateMenu(menu)
	fmt.Println("Menu", menu)
}
