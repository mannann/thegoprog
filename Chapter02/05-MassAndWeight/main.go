package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Scientists measure an object’s mass in kilograms and its weight in newtons. If you know the amount of mass of an object in kilograms, you can calculate its weight in newtons with the following formula:
// weight = mass * 9.8
// Write a program that asks the user to enter an object’s mass, and then calculates its weight.
// If the object weighs more than 500 newtons, display a message indicating that it is too heavy. If the object weighs less than 100 newtons, display a message indicating that it is too light.

func main() {
	fmt.Println("Mass And Weight")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the mass of object : ")
	massString, _ := reader.ReadString('\n')
	massString = strings.TrimSpace(massString)

	mass, _ := strconv.ParseFloat(massString, 64)

	weight := mass * 9.8

	fmt.Printf("Weight of the object is %.2f newtons.\n", weight)

	if weight <= 100 {
		fmt.Println("It's too light")
	} else if weight > 500 {
		fmt.Println("It's too heavy")
	} else {
		fmt.Println("It's ideal weight")
	}

}
