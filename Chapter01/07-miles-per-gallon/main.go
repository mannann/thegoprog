package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A car’s miles-per-gallon (MPG) can be calculated with the following formula:
// MPG = Miles driven Gallons of gas used.
// Write a program that asks the user for the number of miles // driven and the gallons of gas used. /// It should calculate the car’s MPG and display the result

func main() {
	fmt.Println("Miles per Gallon")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the miles travelled : ")
	milesString, _ := reader.ReadString('\n')
	milesString = strings.TrimSpace(milesString)
	miles, _ := strconv.ParseFloat(milesString, 64)

	fmt.Print("Enter the gallons consumed : ")
	gallonString, _ := reader.ReadString('\n')
	gallonString = strings.TrimSpace(gallonString)
	gallons, _ := strconv.ParseFloat(gallonString, 64)

	mpg := miles / gallons

	fmt.Printf("The MPG is : %.3f", mpg)

}
