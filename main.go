package main

import "fmt"

func main() {
	const (
		conferenceName    = "Go Conference"
		conferenceTickets = 50
	)
	var remainingTickets = conferenceTickets
	fmt.Println("Welcome to", conferenceName, "ticket booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
