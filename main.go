package main

import "fmt"

func main() {

	conferenceName := "Go Conference" // you can't use the := operator for const or when explicitly declaring a variable type
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // uint is a type that defines only whole postive integers
	bookings := []string{}         //intializes empty slice(dyname array)

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	// fmt.Scan(&userName)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adderss: ")
	fmt.Scan(&email)

	fmt.Scan("Enter amount of tickets you would like to book: ")
	fmt.Scan(&userTickets)

	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v \n", bookings)

}
