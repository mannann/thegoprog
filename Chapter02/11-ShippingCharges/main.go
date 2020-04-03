package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The Fast Freight Shipping Company charges the following rates:
// 			Weight of Package 												Rate per Pound
// 			2 pounds or less 													$1.50
// 			Over 2 pounds but not more than 6 pounds 	$3.00
// 			Over 6 pounds but not more than 10 pounds $4.00
// 			Over 10 pounds 														$4.75
// Write a program that asks the user to enter the weight of a package and then displays the shipping charges.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the weight of the package in pounds : ")
	weightString, _ := reader.ReadString('\n')
	weightString = strings.TrimSpace(weightString)
	weight, _ := strconv.ParseFloat(weightString, 64)

	if weight >= 10 {
		fmt.Printf("It's cost $4.75 per pound.Total is $%.2f.\n", (weight * 4.75))
	} else if weight >= 6 && weight < 10 {
		fmt.Printf("It's cost $4.00 per pound.Total is $%.2f.\n", (weight * 4.00))
	} else if weight >= 2 && weight < 6 {
		fmt.Printf("It's cost $3.00 per pound.Total is $%.2f.\n", (weight * 3.00))
	} else if weight >= 0 && weight < 2 {
		fmt.Printf("It's cost $1.50 per pound.Total is $%.2f.\n", (weight * 1.50))
	}
}
