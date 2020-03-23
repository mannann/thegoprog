package main

// A customer in a store is purchasing five items. Write a program that asks for the price of
// each item, and then displays the subtotal of the sale, the amount of sales tax, and the total.
// Assume the sales tax is 7 percent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	const taxPercent = 0.07
	fmt.Print("Enter number One: ")
	itemOne, _ := reader.ReadString('\n')
	itemOne = strings.TrimSpace(itemOne)
	one, _ := strconv.ParseFloat(itemOne, 64)

	fmt.Print("Enter number Two: ")
	itemTwo, _ := reader.ReadString('\n')
	itemTwo = strings.TrimSpace(itemTwo)
	two, _ := strconv.ParseFloat(itemTwo, 64)

	fmt.Print("Enter number Three: ")
	itemThree, _ := reader.ReadString('\n')
	itemThree = strings.TrimSpace(itemThree)
	three, _ := strconv.ParseFloat(itemThree, 64)

	fmt.Print("Enter number Four: ")
	itemFour, _ := reader.ReadString('\n')
	itemFour = strings.TrimSpace(itemFour)
	four, _ := strconv.ParseFloat(itemFour, 64)

	fmt.Print("Enter number Five: ")
	itemFive, _ := reader.ReadString('\n')
	itemFive = strings.TrimSpace(itemFive)
	five, _ := strconv.ParseFloat(itemFive, 64)

	total := one + two + three + four + five
	profitAmount := total * taxPercent
	totalAmount := total + profitAmount
	fmt.Printf("The profit amount: $%.3f \n", profitAmount)
	fmt.Printf("The total amount: $%.3f \n", totalAmount)

}
