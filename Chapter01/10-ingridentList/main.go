package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A cookie recipe calls for the following ingredients:
// • 1.5 cups of sugar
// • 1 cup of butter
// • 2.75 cups of flour
// The recipe produces 48 cookies with this amount of the ingredients. Write a program that
// asks the user how many cookies he or she wants to make, and then displays the number of
// cups of each ingredient needed for the specified number of cookies.

func main() {
	fmt.Println("Ingredient Adjuster")
	sugar := 1.5
	butter := 1.0
	flour := 2.75

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the number of Cookies : ")
	cookieString, _ := reader.ReadString('\n')
	cookieString = strings.TrimSpace(cookieString)
	cookie, _ := strconv.ParseFloat(cookieString, 64)

	sugarNeeded := (sugar * cookie) / 48
	butterNeeded := (butter * cookie) / 48
	flourNeeded := (flour * cookie) / 48

	fmt.Printf("Sugar Needed  : %.2f cups.\n", sugarNeeded)
	fmt.Printf("Butter Needed : %.2f cups.\n", butterNeeded)
	fmt.Printf("Flour Needed  : %.2f cups.\n", flourNeeded)
}
