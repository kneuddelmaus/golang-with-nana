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
}
