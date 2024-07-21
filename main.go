package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// slice
	var bookings []string

	// array
	// var bookings [50]string

	// infinite loop
	for {
		var firstName string 
		var lastName string 
		var email string 
		var userTickets uint

		// collect user's information
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 1 && len(lastName) > 1
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber  {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName) 

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are still %v tickets available for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking) // split by space
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of our bookings are: %v\n", firstNames)

			noTicketsAvailable := remainingTickets == 0
			if noTicketsAvailable {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			} else if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @sign")
			} else if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			// continue
		}	
	}
}