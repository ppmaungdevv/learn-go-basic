package main

import "fmt"

func main()  {
	// there is no variable lifting like JS, so makes more sense to declare variable at the top
	conference_name := "GO conference" // this can only be used with var not with const and specific data type declaration 
	const conference_tickets = 50
	var remaining_tickets uint = 50 // uint => unsigned int
	var full_name string
	var user_tickets uint
	var email string
	
	// remaining_tickets = -1 // this will be an error as the varibale was declared as an unsinged int
	
	/* -------- array -----------
	*	are fixed size
	*	var arr = [50]{} <= empty array
	*	var arr = [50]{'m', 'n'} <= with initial value
	*	can accept only the same data type and mixed type won't work
	*	
	*/
	// var bookings [50]string // <= alternative syntac for delcaring empty array

	/* --------- slice --------
	*	is abstraction of an array
	*	more flexible with dynamic size or sub-array of its own
	*	do have a size but resized when needed
	*	to be simple a dynamic size array with different data assigining methods
	*/
	var bookings []string // <= a slice


	fmt.Printf("conference_name is a %T, conference_tickets is an %T, remaining_tickets is an %T\n", conference_name, conference_tickets, remaining_tickets)

	fmt.Printf("Welcome %v booking application\n", conference_name)
	fmt.Printf("Total Seats: %v\nRemaining Tickets: %v\n", conference_tickets, remaining_tickets )
	fmt.Println("Get tickets")

	fmt.Println(&remaining_tickets) // &remainingTickets is a pointer which is a special variable to point the memory location of a variable

	// when using scan, asking for user input, need to use pointer "&var"
	fmt.Println("Enter Full Name:")
	fmt.Scan(&full_name)
	fmt.Println("Enter Email:")
	fmt.Scan(&email)
	fmt.Println("Enter Number of Tickets:")
	fmt.Scan(&user_tickets)
	
	// invalid operation will if the variables in the calculation do not have the same types
	remaining_tickets = remaining_tickets - user_tickets
	// bookings[0] = full_name // use this way to add data to array
	bookings = append(bookings, full_name) // use this way to add data to slice
	// bookings[52] = full_name // out of bond error will show

	fmt.Printf("Array: %v\n", bookings)
	fmt.Printf("First Element: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("%v booked %v tickets. Confirmation mail will be sent to %v\n", full_name, user_tickets, email)
	fmt.Println("Remaining Tickets: ", remaining_tickets)



}
