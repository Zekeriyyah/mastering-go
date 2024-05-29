package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Zekeriyyah/mastering-go/GO-COMPLETE-GUIDE/bank/fileops"
)

const recordFileName = "balance_sheet.txt"

type Account struct {
	Balance float64
}

func main() {
	initialBalance, _ := fileops.GetFloatFromFile(recordFileName)
	// if err != nil {
	// 	fmt.Println("ERROR\t", err)
	// 	return
	// }

	acc := &Account{Balance: initialBalance}

	fmt.Println("Welcome to dummy bank!\nWhat do you want to do?")
	fmt.Println("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit")

	var choice int
	fmt.Print("\n>>>> Your choice: ")
	fmt.Scan(&choice)

	for choice != 4 {
		acc.executeTask(choice)
		fmt.Println("\nWant to perform another task? Press")
		fmt.Println("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit")

		fmt.Print("\n>>>> Your choice: ")
		fmt.Scan(&choice)
	}

}

func (a *Account) executeTask(choice int) {
	switch choice {
	case 1:
		fmt.Println("Your account balance is:", a.Balance)
	case 2:
		var deposit float64
		fmt.Print(">>> Amount to Deposit: $")
		fmt.Scan(&deposit)

		if deposit <= 0 {
			fmt.Println("Invalid amount! Must be greater than 0")
			break
		}
		a.Balance += deposit

		err := fileops.WriteFloatToFile(recordFileName, a.Balance)
		if err != nil {
			fmt.Println("ERROR\t", "transaction failed")
			return
		}

		fmt.Println(deposit, "deposited successfully!")

	case 3:
		var withdrawAmount float64

		fmt.Print(">>> Amount to withdraw: $")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("invalid amount! must be greater than 0")
			break
		}

		if withdrawAmount > a.Balance {
			fmt.Println("sorry! insufficient fund")
			break
		}

		a.Balance -= withdrawAmount
		fmt.Println("TRANSACTION IN PROGRESS....")
		time.Sleep(3 * time.Second)

		err := fileops.WriteFloatToFile(recordFileName, a.Balance)
		if err != nil {
			fmt.Println("ERROR\t", "transaction failed")
			return
		}
		fmt.Println("Withdrawal Successful!\nThank you for banking with us!!!")
	case 4:
		fmt.Println("Thank you for banking with us!!!")
		os.Exit(1)

	default:
		fmt.Println("invalid transaction")
	}
}
