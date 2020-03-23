package main

import "fmt"

// Last month Joe purchased some stock in Acme Software, Inc. Here are the details of the
// purchase:
// • The number of shares that Joe purchased was 2,000.
// • When Joe purchased the stock, he paid $40.00 per share.
// • Joe paid his stockbroker a commission that amounted to 3 percent of the amount he
// paid for the stock.
// Two weeks later Joe sold the stock. Here are the details of the sale:
// • The number of shares that Joe sold was 2,000.
// • He sold the stock for $42.75 per share.
// •  He paid his stockbroker another commission that amounted to 3 percent of the amount
// he received for the stock.
// Write a program that displays the following information:
// • The amount of money Joe paid for the stock.
// • The amount of commission Joe paid his broker when he bought the stock.
// • The amount that Joe sold the stock for.
// • The amount of commission Joe paid his broker when he sold the stock.
// •  Display the amount of money that Joe had left when he sold the stock and paid his
// broker (both times). If this amount is positive, then Joe made a profit. If the amount is
// negative, then Joe lost money.

func main() {
	fmt.Println("Stock Calculator")
	numberOfSharesPurchased := 2000
	numberOfSharesSold := 2000
	purchaseAmount := 40.00
	sellAmount := 42.75
	brokerPercent := 0.03

	amountPaid := float64(numberOfSharesPurchased) * purchaseAmount
	comissionPaidBuy := amountPaid * brokerPercent

	amountSold := float64(numberOfSharesSold) * sellAmount
	comissionPaidSell := amountSold * brokerPercent

	totalProfit := (amountSold + comissionPaidSell) - (amountPaid + comissionPaidBuy)
	fmt.Printf("Amount paid for purchasing : %.3f \n", amountPaid)
	fmt.Printf("Amount paid to broker purchasing : %.3f \n", comissionPaidBuy)
	fmt.Printf("Amount paid for Selling : %.3f \n", amountSold)
	fmt.Printf("Amount paid to broker selling : %.3f \n", comissionPaidSell)
	fmt.Printf("Total Profit : %.3f \n", totalProfit)

}
