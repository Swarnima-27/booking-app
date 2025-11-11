package main

import (
	"fmt"
	"strings"
)

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

	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		userName = firstName + " " + lastName
		fmt.Println("Enter number of tickets to be booked")
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets
		booking = append(booking, userName)
		firstNames := []string{}
		for _, name := range booking {
			fN := strings.Fields(name)
			firstNames = append(firstNames, fN[0])
		}
		fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
		fmt.Printf("First Names of all booked users %v", firstNames)
	}

}
