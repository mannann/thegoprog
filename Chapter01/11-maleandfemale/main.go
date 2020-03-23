package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that asks the user for the number of males and the number of females registered in a class. The program should display the percentage of males and females in the class.
// Hint: Suppose there are 8 males and 12 females in a class. There are 20 students in the
// class. The percentage of males can be calculated as 8 / 20 = 0.4, or 40%. The percentage
// of females can be calculated as 12 / 20 = 0.6, or 60%.

func main() {
	fmt.Println("Male and Female percentage calcukator")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the the number of males students in the class: ")
	maleString, _ := reader.ReadString('\n')
	maleString = strings.TrimSpace(maleString)
	males, _ := strconv.ParseFloat(maleString, 64)
	fmt.Print("Enter the the number of females students in the class: ")
	femaleString, _ := reader.ReadString('\n')
	femaleString = strings.TrimSpace(femaleString)
	females, _ := strconv.ParseFloat(femaleString, 64)

	totalStudents := males + females

	malePercent := (males / totalStudents) * 100
	femalePercent := (females / totalStudents) * 100

	fmt.Printf("Male Percent in the class - %.2f%% \n", malePercent)
	fmt.Printf("Female Percent in the class - %.2f%% \n", femalePercent)
}
