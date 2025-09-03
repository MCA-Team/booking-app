package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets uint = 50 
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var bookings []string

	// Welcome message
	fmt.Printf("Welcome to %v booking application !\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")

	// Ask user for their informations
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your e-mail address: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	// Print the booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v !\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("Total number of bookers: %v\n", bookings)

}