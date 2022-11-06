package main

import "fmt"

func main()  {
	conference_name := "GO conference" // this can only be used with var not with const and specific data type declaration 
	const conference_tickets = 50
	var remaining_tickets uint = 50 // uint => unsigned int

	// remaining_tickets = -1 // this will be an error as the varibale was declared as an unsinged int

	fmt.Printf("conference_name is a %T, conference_tickets is an %T, remaining_tickets is an %T\n", conference_name, conference_tickets, remaining_tickets)

	fmt.Printf("Welcome %v booking application\n", conference_name)
	fmt.Printf("Total Seats: %v\nRemaining Tickets: %v\n", conference_tickets, remaining_tickets )
	fmt.Println("Get tickets")

	fmt.Println(&remaining_tickets) // &remainingTickets is a pointer which is a special variable to point the memory location of a variable

	
	var full_name string
	var user_tickets int
	var email string
	// when using scan, asking for user input, need to use pointer "&var"
	fmt.Println("Enter Full Name:")
	fmt.Scan(&full_name)
	fmt.Println("Enter Email:")
	fmt.Scan(&email)
	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&user_tickets)

	fmt.Printf("%v booked %v tickets. Confirmation mail will be sent to %v\n", full_name, user_tickets, email)

}
