package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-16v ...$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-16v ...$%0.2f \n", "tip: ", b.tip)

	// total
	fs += fmt.Sprintf("%-16v ...$%0.2f", "total: ", total)
	return fs
}

// update the bill's tip
func (b *bill) updateTip(tip float64) {
	// go will automatically copy pointer
	b.tip = tip
}

//  add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save the bill
func (b *bill) save() {
	data := []byte(b.format())

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	error := os.WriteFile(pwd+"/day07/user_input/"+b.name+".txt", data, 0644)
	if error != nil {
		panic(error)
	}

	fmt.Println("Bill saved to file")
}
