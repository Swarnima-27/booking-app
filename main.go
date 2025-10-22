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

	userName = "SWARNIMA"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets", userName, userTickets)

}
