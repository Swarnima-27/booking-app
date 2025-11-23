package main

import (
	"fmt"
	"strings"
)

func main() {

	var eventName = "Jalsa"
	const totalTickets = 50
	var remainingTickets uint = 50

	greetGuests(eventName, totalTickets, remainingTickets)

	var userName string
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var booking []string

	for {
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		userName = firstName + " " + lastName
		fmt.Println("Enter email ID")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets to be booked")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets

			booking = append(booking, userName)
			firstNames := []string{}
			for _, name := range booking {
				fN := strings.Fields(name)
				firstNames = append(firstNames, fN[0])
			}
			fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
			fmt.Printf("First Names of all booked users %v \n", firstNames)
			fmt.Printf("Tickets Remaining %v \n", remainingTickets)
		} else {
			fmt.Println("ERROR IN BOOKING TICKETS!!!!")
			if !isValidName {
				fmt.Println("Name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Incorrect mail, missing valid format")
			}
			if !isValidTicket {
				fmt.Printf("Invalid number of tickets entered")
			}
		}
	}

}
func greetGuests(eventName string, totalTickets int, remainingTickets uint) {
	fmt.Println("Welcome to", eventName, "Booking Application")
	fmt.Printf("We have total of %v tickets and %v tickets available for booking\n", totalTickets, remainingTickets)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
