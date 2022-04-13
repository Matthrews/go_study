package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":    9.45,
		"pie":     7.00,
		"vinegar": 8.54,
		"coffee":  32.00,
	}

	// map key will be alphabetical
	fmt.Println("Menu", menu)
	fmt.Println("Menu Item", menu["pie"])

	// loop maps
	// original order
	for k, v := range menu {
		fmt.Println("K-V", k, v)
	}

	phonebook := map[int]string{
		12432432: "margio",
		98783453: "peach",
		65465343: "luigi",
	}

	// map key will be alphabetical
	fmt.Println("PhoneBook", phonebook)
	fmt.Println("PhoneBook Item", phonebook[65465343])
}
