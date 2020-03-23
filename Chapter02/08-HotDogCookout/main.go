package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Assume that hot dogs come in packages of 10, and hot dog buns come in packages of 8. Write a program that calculates the number of packages of hot dogs and the number of packages of hot dog buns needed for a cookout, with the minimum amount of leftovers.
// The program should ask the user for the number of people attending the cookout and the number of hot dogs each person will be given. The program should display the following details:
// •	 The minimum number of packages of hot dogs required
// •	 The minimum number of packages of hot dog buns required
// •	 The number of hot dogs that will be left over
// •	 The number of hot dog buns that will be left over

func main() {
	fmt.Println("Hot dog cookout")
	reader := bufio.NewReader(os.Stdin)
	var hotDog = 10
	var hotDogBuns = 8

	fmt.Print("Enter the number of persons attending : ")
	personString, _ := reader.ReadString('\n')
	personString = strings.TrimSpace(personString)
	persons, _ := strconv.Atoi(personString)
	fmt.Print("Enter the number of hot dogs per person : ")
	bunsPerString, _ := reader.ReadString('\n')
	bunsPerString = strings.TrimSpace(bunsPerString)
	buns, _ := strconv.Atoi(bunsPerString)

	minimumPackagesOfDogs := (persons * buns) / hotDog
	minimumPackagesOfBuns := (persons * buns) / hotDogBuns

	leftDogs := persons % hotDog
	leftBuns := persons % hotDogBuns

	fmt.Println("Hot dog packets required : ", minimumPackagesOfDogs)
	fmt.Println("Hot dog bun packets required : ", minimumPackagesOfBuns)
	fmt.Println("Hot dogs left : ", leftDogs)
	fmt.Println("Hot dogs buns left: ", leftBuns)

}
