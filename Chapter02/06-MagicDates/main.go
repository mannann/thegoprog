package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The date June 10, 1960, is special because when it is written in the following format, the month times the day equals the year: 6/10/60
// Design a program that asks the user to enter a month (in numeric form), a day, and a twodigit year. The program should then determine whether the month times the day equals the year. If so, it should display a message saying the date is magic. Otherwise, it should display a message saying the date is not magic.

func main() {
	fmt.Println("The magic date")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the date in DD: ")
	dateString, _ := reader.ReadString('\n')
	dateString = strings.TrimSpace(dateString)
	date, _ := strconv.ParseFloat(dateString, 64)
	fmt.Print("Enter the month in MM: ")
	monthString, _ := reader.ReadString('\n')
	monthString = strings.TrimSpace(monthString)
	month, _ := strconv.ParseFloat(monthString, 64)

	fmt.Print("Enter the year in YY: ")
	yearString, _ := reader.ReadString('\n')
	yearString = strings.TrimSpace(yearString)
	year, _ := strconv.ParseFloat(yearString, 64)

	if date*month == year {
		fmt.Println("It's a magic date")
	} else {
		fmt.Println("Better luck with new date and month")
	}
}
