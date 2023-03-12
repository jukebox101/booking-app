package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference" // you can't use the := operator for const or when explicitly declaring a variable type
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // uint is a type that defines only whole postive integers
	bookings := []string{}         //intializes empty slice(dynamic array)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		fmt.Println("Enter amount of tickets you would like to book: ")
		fmt.Scan(&userTickets)

		if remainingTickets <= userTickets {
			fmt.Printf("We only have %v tickets left, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue // causes loop to skip the rest of its body, and retests the condition
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		// _ is used to tell compiler there is a variable that is not explicitly being used
		// for index, booking := range bookings { }
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of our bookings: %v \n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}

}
