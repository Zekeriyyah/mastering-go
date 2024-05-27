package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var yearOfReturn = 10

	fmt.Print(">>>Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print(">>>Expected Rate of Return: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print(">>>Year: ")
	fmt.Scan(&yearOfReturn)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(yearOfReturn))
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, float64(yearOfReturn))

	fmt.Printf("Future Value: %v\n", futureValue)
	fmt.Printf("Future value under inflation: %v\n", futureRealValue)

}
