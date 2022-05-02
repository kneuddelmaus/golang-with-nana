package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint8
}

func main() {
	const (
		conferenceName    string = "Go Conference"
		conferenceTickets uint8  = 50
	)
	var remainingTickets uint8 = conferenceTickets
	fmt.Println("Welcome to", conferenceName, "ticket booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	var bookings []UserData
	var wg = sync.WaitGroup{}

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTicket := validateInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTicket {
			bookTickets(firstName, lastName, email, userTickets, bookings, &remainingTickets)
			fmt.Printf("Thank you %v %v for purchasing %v tickets for %v\n", firstName, lastName, userTickets, conferenceName)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
			wg.Add(1)
			go emailTickets(userTickets, email, &wg)
		} else {
			errorPrinter(isValidName, isValidEmail, isValidUserTicket)
		}
	}
	wg.Wait()
}

func getUserInput() (string, string, string, uint8) {
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

	return firstName, lastName, email, userTickets
}

func validateInputs(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicket
}

func bookTickets(firstName string, lastName string, email string, userTickets uint8, bookings []UserData, remainingTickets *uint8) {
	var user = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, user)
	*remainingTickets -= userTickets
}

func errorPrinter(isValidName bool, isValidEmail bool, isValidUserTicket bool) {
	if !isValidName {
		fmt.Println("Invalid first name or last name provided")
	}
	if !isValidEmail {
		fmt.Println("Invalid email provided")
	}
	if !isValidUserTicket {
		fmt.Println("Invalid ticket number requested")
	}
}

func emailTickets(userTickets uint8, email string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Sending %v tickets to %v\n", userTickets, email)
	time.Sleep(60 * time.Second)
	fmt.Printf("Email sent to %v\n", email)
}
