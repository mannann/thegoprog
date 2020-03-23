package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that calculates the total amount of a meal purchased at a restaurant. The
// program should ask the user to enter the charge for the food, and then calculate the amount
// of a 18 percent tip and 7 percent sales tax. Display each of these amounts and the total.

func main() {
	fmt.Println("The BILL")
	reader := bufio.NewReader(os.Stdin)
	const tip float64 = 0.18
	const salesTax float64 = 0.07
	fmt.Print("Enter the meal billamount: ")
	totalString, _ := reader.ReadString('\n')
	totalString = strings.TrimSpace(totalString)
	total, _ := strconv.ParseFloat(totalString, 64)

	totalAmount := total + (total * tip) + (total * salesTax)

	fmt.Printf("Meal Bill $%.2f\n", total)
	fmt.Printf("Meal Tip  %.2f \n", tip)
	fmt.Printf("Meal Tax  %.2f \n", salesTax)
	fmt.Println("---------------------------")
	fmt.Printf("Total Bill $%.3f\n", totalAmount)

}
