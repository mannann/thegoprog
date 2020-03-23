package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The colors red, blue, and yellow are known as the primary colors because they cannot be made by mixing other colors. When you mix two primary colors, you get a secondary color, as shown here:
// - When you mix red and blue, you get purple.
// - When you mix red and yellow, you get orange.
// - When you mix blue and yellow, you get green.
// Design a program that prompts the user to enter the names of two primary colors to mix. If the user enters anything other than “red,” “blue,” or “yellow,” the program should display an error message. Otherwise, the program should display the name of the secondary color that results

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Pick two colors from the list 'RED', 'BLUE', 'YELLOW' in Caps.")
	fmt.Print("Enter color - 1: ")
	colorOne, _ := reader.ReadString('\n')
	colorOne = strings.ToLower(strings.TrimSpace(colorOne))
	fmt.Println(colorOne)
	fmt.Print("Enter color - 2: ")
	colorTwo, _ := reader.ReadString('\n')
	colorTwo = strings.ToLower(strings.TrimSpace(colorTwo))
	fmt.Println(colorTwo)
	if colorOne == "red" && colorTwo == "blue" {
		fmt.Println("You got PURPLE")
	} else if colorOne == "red" && colorTwo == "yellow" {
		fmt.Println("You got ORANGE")
	} else if colorOne == "blue" && colorTwo == "yellow" {
		fmt.Println("You got GREEN")
	} else {
		fmt.Println("Invalid Choice")
	}
}
