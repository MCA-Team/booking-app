package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets uint = 50 
	var remainingTickets uint = 50
	var userName string
	var userTickets uint

	fmt.Printf("Welcome to %v booking application !\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")


	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}