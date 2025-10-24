package main

import "fmt"

func main() {

	var eventName = "Jalsa"
	const totalTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", eventName, "Booking Application")
	fmt.Printf("We have total of %v tickets and %v tickets available for booking\n", totalTickets, remainingTickets)

	var userName string
	var userTickets int

	fmt.Println("Enter your name")
	fmt.Scan(&userName)
	fmt.Println("Enter number of tickets to be booked")
	fmt.Scan(&userTickets)
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
