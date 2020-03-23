package main

import "fmt"

// Assuming there are no accidents or delays, the distance that a car travels down the interstate can be calculated with the following formula:
// Distance = speec * time
// A car is traveling at 70 miles per hour. Write a program that displays the following:
// 	- The distance the car will travel in 6 hours
// 	- The distance the car will travel in 10 hours
// 	- The distance the car will travel in 15 hours

func main() {
	fmt.Println("Travelled distance")

	speed := 70.00

	inSix := speed * 6
	fmt.Printf("Distance Travelled in 06 hours : %.3f \n", inSix)

	inTen := speed * 10
	fmt.Printf("Distance Travelled in 10 hours : %.3f \n", inTen)

	inFifteen := speed * 15
	fmt.Printf("Distance Travelled in 15 hours : %.3f \n", inFifteen)

}
