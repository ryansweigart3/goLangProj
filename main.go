package main

import (
	"fmt"
	"sync"
	"time"
)

var confName = "Cat Conference"

const confTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidNumber {

			// book tickets
			bookTickets(userTickets, firstName, lastName, confName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			//print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Oh no! Our conference is all booked. Please come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address is incorrect format.")
			}
			if !isValidNumber {
				fmt.Println("Please enter a valid number of tickets.")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", confName)
	fmt.Printf("Cats have 9 lives - we have %v tickets! Currently %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, confName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create map for user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is: %v\n", userData)

	fmt.Printf("Thank you, %v %v for booking %v tickets to the %v! You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, confName, email)
	fmt.Printf("%v tickets remain for %v\n", remainingTickets, confName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("######")
}
