package main

import "fmt"

func main() {
	const totalTickets = 10
	//var remainingTickets int
	var name string
	var requiredTickets int
	var email string
	var confirm bool
	fmt.Printf("Welcome to Example-Ticketing system and we have total %v tickets.\n", totalTickets)
	fmt.Println("May I have your name please?")
	fmt.Scan(&name)
	fmt.Println("How many tickets do you want?")
	fmt.Scan(&requiredTickets)
	fmt.Println("Can I can get your email to confirm your purchase?")
	fmt.Scan(&email)
	fmt.Printf("Alright! Please confirm me with 'yes' if I am right or you can say 'no' to repeat.\nHere is the details I have. Your name is %v, You need %v tickets and your email is %v.\n", name, requiredTickets, email)
	fmt.Scan(&confirm)
}
