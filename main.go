package main

import "fmt"

func main() {
	const (
		conferenceName    string = "Go Conference"
		conferenceTickets uint8  = 50
	)
	var remainingTickets uint8 = conferenceTickets
	fmt.Println("Welcome to", conferenceName, "ticket booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var (
		firstName   string
		lastName    string
		email       string
		userTickets uint8
	)

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you intend to purchase")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for purchasing %v tickets for %v\n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}
