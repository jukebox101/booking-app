package main

import "fmt"

func main() {

	conferenceName := "Go Conference" // you can't use the := operator for const or when explicitly declaring a variable type
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // uint is a type that defines only whole postive integers

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
