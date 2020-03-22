package main

import (
	"fmt"
)

// A company has determined that its annual profit is typically 23 percent of total sales. Write a program that asks the user to enter the projected amount of total sales, and then displays the profit that will be made from that amount.
// Hint: Use the value 0.23 to represent 23 percent.

func main() {
	// Sales Prediction
	const tax float64 = 0.23
	var saleAmount float64
	fmt.Print("Enter the sale amount: ")
	fmt.Scan(&saleAmount)
	profit := saleAmount * tax
	total := saleAmount + profit
	fmt.Printf("Total amount after profit: $%f", total)

}
