package main

import "fmt"

func main() {
	myBill := newBill("mario's bill")

	myBill.updateTip(10.00)

	myBill.addItem("peer", 5.68)

	formattedBill := myBill.format()

	fmt.Println(formattedBill)
}

// go run day07/struct/main.go day07/struct/bill.go
