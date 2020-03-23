package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that will ask the user to enter the amount of a purchase. The program
// should then compute the state and county sales tax. Assume the state sales tax is 5 percent
// and the county sales tax is 2.5 percent. The program should display the amount of the purchase,
// the state sales tax, the county sales tax, the total sales tax, and the total of the sale
// (which is the sum of the amount of purchase plus the total sales tax).
// Hint: Use the value 0.025 to represent 2.5 percent, and 0.05 to represent 5 percent.

func main() {
	fmt.Println("Sales tax program")
	reader := bufio.NewReader(os.Stdin)

	const stateTax = 0.05
	const countryTax = 0.025
	fmt.Print("Enter the amount of purchase: ")
	amount, _ := reader.ReadString('\n')
	amount = strings.TrimSpace(amount)
	purchaseAmount, _ := strconv.ParseFloat(amount, 64)
	totalAmount := purchaseAmount * 0.05 * 0.025

	fmt.Printf("Total state tax   	: %.3f \n", stateTax)
	fmt.Printf("Total country tax 	: %.3f \n", countryTax)
	fmt.Printf("Total Amount      	: %.3f \n", totalAmount)
}
