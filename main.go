package main

import (
	"booking-ticket/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]User, 0)

var wg = sync.WaitGroup{}

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUser()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidEmail && isValidName && isValidTicketCount {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			printFirstNames()
			if remainingTickets == 0 {
				fmt.Println("\nOur conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name is too short <2.")
			}
			if !isValidEmail {
				fmt.Println("Email address is not valid.")
			}
			if !isValidTicketCount {
				fmt.Println("Number of tickets is invalid.")
			}
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceName, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("\nEnter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("\nEnter your email name: ")
	fmt.Scan(&email)

	fmt.Println("\nEnter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("Current bookings: %v\n", firstNames)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	var userData = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("\nUser %v booked %v tickets\n. You will receive a confirmation email.\n", firstName, userTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("\n##############################################")
	fmt.Printf("\n%v %v booked %v tickets\nThank you.\n", firstName, lastName, userTickets)
	fmt.Println("##############################################")
	wg.Done()
}
