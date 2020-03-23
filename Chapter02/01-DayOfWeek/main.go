package main

import (
	"fmt"
)

// Write a program that asks the user for a number in the range of 1 through 7.
// The program should display the corresponding day of the week, where 1 = Monday, 2 = Tuesday,
// 3 = Wednesday, 4 = Thursday, 5 = Friday, 6 = Saturday, and 7 = Sunday.
// The program should display an error message if the user enters a number that is outside the range of 1 through 7.

func main() {
	fmt.Println("Days of the week")

	// reader := bufio.NewReader(os.Stdin)
	var number int

	fmt.Print("Enter the day number >> ")

	fmt.Scan(&number)

	if number == 1 {
		fmt.Println("Monday")
	} else if number == 2 {
		fmt.Println("Tuesday")
	} else if number == 3 {
		fmt.Println("Wednesday")
	} else if number == 4 {
		fmt.Println("Thursday")
	} else if number == 5 {
		fmt.Println("Friday")
	} else if number == 6 {
		fmt.Println("Saturday")
	} else if number == 7 {
		fmt.Println("Sunday")
	} else {
		fmt.Println("Wrong Input")
	}
}
