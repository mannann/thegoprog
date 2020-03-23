package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that asks the user to enter a person’s age. The program should display a message indicating whether the person is an infant, a child, a teenager, or an adult.
// Following are the guidelines:
// - If the person is 1 year old or less, he or she is an infant.
// - If the person is older than 1 year, but younger than 13 years, he or she is a child.
// - If the person is at least 13 years old, but less than 20 years old, he or she is a teenager.
// - If the person is at least 20 years old, he or she is an adult.

func main() {
	fmt.Println("Age classifier")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the age : ")
	ageString, _ := reader.ReadString('\n')
	ageString = strings.TrimSpace(ageString)

	age, _ := strconv.ParseFloat(ageString, 64)

	if age >= 0 && age < 1 {
		fmt.Println("She/He is an Infant")
	} else if age >= 1 && age <= 13 {
		fmt.Println("She/He is an yonger")
	} else if age >= 20 {
		fmt.Println("She/He is an adult")
	}

}
