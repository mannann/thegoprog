package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// On a roulette wheel, the pockets are numbered from 0 to 36. The colors of the pockets are as follows:
// •	 Pocket 0 is green.
// •	 For pockets 1 through 10, the odd-numbered pockets are red and the even-numbered pockets are black.
// •	 For pockets 11 through 18, the odd-numbered pockets are black and the even-numbered pockets are red.
// •	 For pockets 19 through 28, the odd-numbered pockets are red and the even-numbered pockets are black.
// •	 For pockets 29 through 36, the odd-numbered pockets are black and the even-numbered pockets are red.
// Write a program that asks the user to enter a pocket number and displays whether the pocket is green, red, or black. The program should display an error message if the user enters a number that is outside the range of 0 through 36.

func main() {
	fmt.Println("The Roulette Wheel")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number : ")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, _ := strconv.Atoi(numStr)

	if num < 0 && num > 36 {
		fmt.Println("Invalid Input")
	} else {
		if num == 0 {
			fmt.Println("Number 0 is Green")
		} else if num >= 1 && num <= 10 {
			if num%2 == 0 {
				fmt.Printf("Number %d is Black.\n", num)
			} else {
				fmt.Printf("Number %d is Red.\n", num)
			}
		} else if num >= 11 && num <= 18 {
			if num%2 == 0 {
				fmt.Printf("Number %d is Red.\n", num)
			} else {
				fmt.Printf("Number %d is Black.\n", num)
			}
		} else if num >= 19 && num <= 28 {
			if num%2 == 0 {
				fmt.Printf("Number %d is Black.\n", num)
			} else {
				fmt.Printf("Number %d is Red.\n", num)
			}
		} else if num >= 29 && num <= 36 {
			if num%2 == 0 {
				fmt.Printf("Number %d is Red.\n", num)
			} else {
				fmt.Printf("Number %d is Black.\n", num)
			}
		}
	}

}
