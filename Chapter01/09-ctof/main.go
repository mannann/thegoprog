package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a program that converts Celsius temperatures to Fahrenheit temperatures. The formula
// is as follows:
// F = (9/5)*c + 32
// The program should ask the user to enter a temperature in Celsius, and then display the
// temperature converted to Fahrenheit.

func main() {
	fmt.Println("Temperature Converter")
	var celcius float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the temperature in celcius : ")
	celciusString, _ := reader.ReadString('\n')
	celciusString = strings.TrimSpace(celciusString)
	celcius, _ = strconv.ParseFloat(celciusString, 64)

	fahrenheit := (1.8 * celcius) + 32

	fmt.Printf("%.2f C = %.2f \n", celcius, fahrenheit)

}
