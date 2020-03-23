package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that prompts the user to enter a number within the range of 1 through 10.
// The program should display the Roman numeral version of that number. If the number is outside the range of 1 through 10, the program should display an error message. The following table shows the Roman numerals for the numbers 1 through 10:
// Number Roman Numeral
// 1 I
// 2 II
// 3 III
// 4 IV
// 5 V
// 6 VI
// 7 VII
// 8 VIII
// 9 IX
// 10 X

// We can write this program using if conditions or switch statements

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number : ")

	numString, _ := reader.ReadString('\n')
	numString = strings.TrimSpace(numString)
	num, _ := strconv.ParseFloat(numString, 64)

	switch {
	case num == 1:
		fmt.Println("I")
	case num == 2:
		fmt.Println("II")
	case num == 3:
		fmt.Println("III")
	case num == 4:
		fmt.Println("IV")
	case num == 5:
		fmt.Println("V")
	case num == 6:
		fmt.Println("XI")
	case num == 7:
		fmt.Println("XII")
	case num == 8:
		fmt.Println("XIII")
	case num == 9:
		fmt.Println("IX")
	case num == 10:
		fmt.Println("X")
	default:
		fmt.Println("Invalid Input")
	}

}
