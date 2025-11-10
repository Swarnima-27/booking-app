package main

import "fmt"

func main() {

	var eventName = "Jalsa"
	const totalTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", eventName, "Booking Application")
	fmt.Printf("We have total of %v tickets and %v tickets available for booking\n", totalTickets, remainingTickets)

	var userName string
	var firstName string
	var lastName string
	var userTickets int
	var booking []string

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	userName = firstName + " " + lastName
	fmt.Println("Enter number of tickets to be booked")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	booking = append(booking, userName)
	fmt.Printf("The whole slice %v\n", booking)
	fmt.Printf("First value %v\n", booking[0])
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
