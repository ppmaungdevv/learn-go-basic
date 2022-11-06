package main

import "fmt"

func main()  {
	conferenceName := "GO conference" // this can only be used with var not with const and specific data type declaration 
	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint => unsigned int

	// remainingTickets = -1 // this will be an error as the varibale was declared as an unsinged int

	fmt.Printf("conferenceName is a %T, conferenceTickets is an %T, remainingTickets is an %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome %v booking application\n", conferenceName)
	fmt.Printf("Total Seats: %v\nRemaining Tickets: %v\n", conferenceTickets, remainingTickets )
	fmt.Println("Get tickets")

	var useName string
	var userTickets int
	useName = "Ma"
	userTickets = 2
	fmt.Printf("%v booked %v tickets\n", useName, userTickets)

}
