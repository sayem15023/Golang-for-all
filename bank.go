package main

import "fmt"

func bankAddMoeny(balance *int, amount int) {
	fmt.Println("Add amout to main balance", amount, "to a balance")
	fmt.Println("Memory Address of balance in addMoney:", balance)
	*balance += amount
	fmt.Println("new balance inside add money", *balance)

}

func witdrawMoney(balance *int, amount int) {

}
