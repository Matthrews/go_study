package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createbill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Create a new bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return newBill(name)
}

func main() {
	myBill := createbill()

	// myBill.updateTip(10.00)

	// myBill.addItem("peer", 5.68)

	// formattedBill := myBill.format()

	fmt.Println(myBill)
}

// go run day07/user_input/main.go day07/user_input/bill.go
