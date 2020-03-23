// One acre of land is equivalent to 43,560 square feet. Write a program that asks the user to enter the total square feet in a tract of land and calculates the number of acres in the tract.
// Hint: Divide the amount entered by 43,560 to get the number of acres.package main

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("The Acre Program")
	const oneAcre = 43560.0
	var theFeet float64
	fmt.Print("Enter the number of square feet: ")
	feet, _ := reader.ReadString('\n')
	feet = strings.TrimSpace(feet)
	theFeet, _ = strconv.ParseFloat(feet, 64)
	numberOfAcres := theFeet / oneAcre
	fmt.Printf("Total Number of Acres: %v\n", numberOfAcres)
}
