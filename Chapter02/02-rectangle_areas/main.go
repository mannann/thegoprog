package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The area of a rectangle is the rectangle’s length times its width. Write a program that asks for the length and width of two rectangles. The program should tell the user which rectangle has the greater area, or if the areas are the same.

func main() {
	fmt.Println("Area of rectangles")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Length of traingle A : ")
	rectALString, _ := reader.ReadString('\n')
	rectALString = strings.TrimSpace(rectALString)
	reactAL, _ := strconv.ParseFloat(rectALString, 64)
	fmt.Print("Width of traingle A : ")

	rectAWString, _ := reader.ReadString('\n')
	rectAWString = strings.TrimSpace(rectAWString)
	reactAW, _ := strconv.ParseFloat(rectAWString, 64)

	fmt.Print("Length of traingle B : ")

	rectBLString, _ := reader.ReadString('\n')
	rectBLString = strings.TrimSpace(rectBLString)
	reactBL, _ := strconv.ParseFloat(rectBLString, 64)
	fmt.Print("Width of traingle B : ")

	rectBWString, _ := reader.ReadString('\n')
	rectBWString = strings.TrimSpace(rectBWString)
	reactBW, _ := strconv.ParseFloat(rectBWString, 64)

	areaTraingleA := reactAL * reactAW
	areaTraingleB := reactBL * reactBW

	if areaTraingleA > areaTraingleB {
		fmt.Println("Traingle A has larger area than B")
	} else if areaTraingleB > areaTraingleA {
		fmt.Println("Traingle B has larger area than A")
	} else {
		fmt.Println("Both traingles are same in area")
	}
}
