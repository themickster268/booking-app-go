package main

import (
	"booking-app-go/helper"
	"fmt"
	"time"
	"sync"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // OR var bookings = []string{} OR bookings := []string{}

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for - infinite loop
	// for ;i < 10; - while loop
	// for i := 0; i < 10; i++ - for loop

	// for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				// break out of indefinite loop
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("The First name or Last name you have entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The Email address you have entered does not contain an '@' sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you have entered is invalid.")
			}
		}
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var firstName = booking.firstName
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	firstNames := getFirstNames()

	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}


func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}