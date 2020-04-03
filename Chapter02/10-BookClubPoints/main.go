package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Serendipity Booksellers has a book club that awards points to its customers based on the number of books purchased each month. The points are awarded as follows:
// •	 If a customer purchases 0 books, he or she earns 0 points.
// •	 If a customer purchases 2 books, he or she earns 5 points.
// •	 If a customer purchases 4 books, he or she earns 15 points.
// •	 If a customer purchases 6 books, he or she earns 30 points.
// •	 If a customer purchases 8 or more books, he or she earns 60 points.
// Write a program that asks the user to enter the number of books that he or she has purchased this month and displays the number of points awarded.

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the book count : ")
	countString, _ := reader.ReadString('\n')
	countString = strings.TrimSpace(countString)

	count, _ := strconv.Atoi(countString)

	if count == 0 {
		fmt.Println("You have 0 points")
	} else if count >= 8 {
		fmt.Println("You have 60 points")
	} else if count >= 6 && count < 8 {
		fmt.Println("You have 30 points")
	} else if count >= 4 && count < 6 {
		fmt.Println("You have 15 points")
	} else if count >= 2 && count < 4 {
		fmt.Println("You have 5 points")
	}
}
