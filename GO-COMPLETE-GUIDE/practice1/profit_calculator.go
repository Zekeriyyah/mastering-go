package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print(">>> Revenue: $")
	fmt.Scan(&revenue)

	fmt.Print(">>> Expenses: $")
	fmt.Scan(&expenses)

	fmt.Print(">>> Tax Rate: $")
	fmt.Scan(&taxRate)

	var earningsBeforeTax = revenue - expenses
	var profit = earningsBeforeTax * (1 - taxRate/100)

	fmt.Printf("Earnings Before Tax: $%.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("EBT/PROFIT: %.2f\n", earningsBeforeTax/profit)

}
