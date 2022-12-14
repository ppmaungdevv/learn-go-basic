package main

import (
	"fmt"
	"time"
)

/* 
* --------- interface ----------
*	a collection struct behaviours, simplest and easiest way to say about interface
*	for more understanding, watch this => https://www.youtube.com/watch?v=lh_Uv2imp14&list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q&index=22
*/

type User interface {
	PrintName()
	PrintInfo()
}

type Person struct {
	name, email string // variable with the same data type can be written in comma separated line
	Dob time.Time
	location string
}

/* 
*	--------- struct methods ------------
*	is similar to function but include a receiver
*	specify the method receiver between the func keyword and the name of the function
*	func (variable Struct) FuncName() => (variable Struct) -> is a receiver
*/

func (p Person) PrintName() {
	fmt.Printf("Person Name: %v \n", p.name)
}

func (p Person) PrintInfo() {
	fmt.Printf("Person Info: %v \n", p)
}

/* 
*	-------- struct method with pointer receiver ------------
*	to modify the data of a receiver from the method, the receiver must be a pointer
*	to call a pointer receiver method, create the Person type instance by providing the ampersand (&) operator
*	If the struct has a pointer receiver on some its methods, it is better to use it for the rest of the methods because it enables better consistency and predictability for the struct behaviors
*/

// func (p *Person) PrintName() {
// 	fmt.Printf("Person Name %v \n", p.name)
// }

// func (p *Person) PrintInfo() {
// 	fmt.Printf("Person Info %v \n", p)
// }

// func (p *Person) ChangeLocation(new_location string) {
// 	fmt.Printf("OG localition %v \n", p.location)
// 	p.location = new_location
// 	fmt.Printf("Changed location %v \n", p.location)
// }


/* 
*	------------- type embedding for composition -------------
*	create bigger and more concrete types by embedding the Person type
*	can override the embedded methods by using the same method name
*/

type Admin struct {
	Person
	roles []string
}

// func (a Admin) PrintAdminInfo()  { // <= this won't override the embedded method
func (a Admin) PrintInfo()  {
	a.Person.PrintInfo()
	fmt.Printf("Roles: ")
	for _, v := range a.roles {
		fmt.Printf(v)
	}
	fmt.Printf("\n")
}

type Member struct {
	Person
	skills []string
}

// func (m Member) PrintMemberInfo()  { // <= this won't override the embedded method
func (m Member) PrintInfo()  {
	m.Person.PrintInfo()
	fmt.Printf("Skills: ")
	for _, v := range m.skills {
		fmt.Printf(" %v", v)
	}
	fmt.Printf("\n")
}

type Team struct {
	name, description string
	Users []User
}

func (t Team) PrintTeamInfo() {
	fmt.Printf("Team: %s \nDesc: %s \n", t.name, t.description)
	for _, user := range t.Users {
		user.PrintName()
		user.PrintInfo()
	}
}

func main()  {
	p :=  Person {
		name: "Pyae",
		email: "pp@m.co",
		Dob: time.Date(1990, time.August, 23, 0, 0, 0, 0, time.UTC),
		location: "Lashio",
	}

	a := Admin {
		Person: p,
		roles: []string{"Senior"},
	}

	m := Member{
		Person: Person {
			name: "Phyo",
			email: "phyo@m.co",
			Dob: time.Date(1991, time.August, 23, 0, 0, 0, 0, time.UTC),
			location: "Lashio",
		},
		skills: []string{"Vue", "Nuxt", "GO", "AWS"},
	}

	u := []User{a, m}

	team := Team{
		name: "GO",
		description: "GO learn team",
		Users: u,
	}

	team.PrintTeamInfo()

	// for _, user := range u {
	// 	user.PrintName()
	// 	user.PrintInfo()
	// }

	// fmt.Printf("Person type %v \n", p)
	// p.PrintName()
	// p.PrintInfo()
	// // as both Member and Admin embedded the Person, all the methods of Person are available to them
	// a.PrintName()
	// a.PrintInfo()
	// // a.PrintAdminInfo()
	// m.PrintName()
	// m.PrintInfo()
	// // m.PrintMemberInfo()
}

// import (
// 	"Booking-app/validator" // when using another package, need to add module name of that package
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// var conference_name = "GO conference" // cannot use := declaration in package level varibles
// const conference_tickets = 50

// var remaining_tickets uint = 50

// /*
// *	----------- struct (stands for structure)
// *	allows to create mixed data types of data structure, a little similar to objects in JS
// *	"type" keyword needs to create struct
// *	type StructName struct{}
// *
// *	type keywords create a new data type with the name behind it
//  */
// type UserInfoStruct struct {
// 	name    string
// 	email   string
// 	age     uint
// 	married bool
// }

// func main() {
// 	// there is no variable lifting like JS, so makes more sense to declare variable at the top
// 	// conference_name := "GO conference" // this can only be used with var not with const and specific data type declaration
// 	// const conference_tickets = 50
// 	// var remaining_tickets uint = 50 // uint => unsigned int

// 	// remaining_tickets = -1 // this will be an error as the varibale was declared as an unsinged int

// 	// assigning to a struct data type variable
// 	var user_info = UserInfoStruct{
// 		name:    "Pyae",
// 		email:   "p@m.co",
// 		age:     27,
// 		married: false,
// 	}

// 	fmt.Printf("user info struct: %v \n", user_info)
// 	fmt.Printf("get name from user info struct: %v \n", user_info.name)

// 	greetUser()

// 	/* -------- array -----------
// 	*	are fixed size
// 	*	var arr = [50]{} <= empty array
// 	*	var arr = [50]{'m', 'n'} <= with initial value
// 	*	can accept only the same data type and mixed type won't work
// 	*
// 	 */
// 	// var bookings [50]string // <= alternative syntac for delcaring empty array

// 	/* --------- slice --------
// 		*	is abstraction of an array
// 		*	more flexible with dynamic size or sub-array of its own
// 		*	do have a size but resized when needed
// 		*	to be simple a dynamic size array with different data assigining methods
// 		*	read this article for detail understanding about slice => https://gosamples.dev/capacity-and-length/
// 	    *	a slice is a data structure describing a piece of an array with three properties (ptr, len, cap)
// 	    *	var slice_test = make([]int, 2, 3)
// 	    *	var slice_test = make([]int, 3)
// 	    *	slice_test[0] = 10
// 		*	slice_test[1] = 20
// 		*	fmt.Printf("last index %v \n", slice_test[2])
// 	    *	slice_test = append(slice_test, 20)
// 	    *	slice_test = append(slice_test, 30)
// 	    *	slice_test = append(slice_test, 40)
// 	    *	slice_test = append(slice_test, 50)
// 	    *	slice_test = append(slice_test, 60)
// 	    *	slice_test = append(slice_test, 70)
// 	    *	slice_test = append(slice_test, 80)
// 	    *	slice_test = append(slice_test, 90)
// 	    *	slice_test = append(slice_test, 100)
// 	    *	slice_test = append(slice_test, 110)

// 	    *	fmt.Printf("slice test %v \n", slice_test)
// 	    *	fmt.Printf("slice capacity %v \n", cap(slice_test))
// 	    *	fmt.Printf("slice length %v \n", len(slice_test))

// 	    *	var a []int            // nil slice
// 	    *	b := []int{0, 1, 2, 3} // slice initialized with specified array
// 	    *	c := make([]int, 4)    // slice of size 4 initialized with zero-valued array of size 4
// 	    *	d := make([]int, 4, 5) // slice of size 4 initialized with zero-valued array of size 5

// 	    *	fmt.Printf("a: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(a), cap(a), a, a, a == nil)
// 	    *	fmt.Printf("b: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(b), cap(b), b, b, b == nil)
// 	    *	fmt.Printf("c: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(c), cap(c), c, c, c == nil)
// 	    *	fmt.Printf("d: length: %d, capacity: %d, pointer to underlying array: %p, data: %v, is nil: %t\n", len(d), cap(d), d, d, d == nil)
// 	*/
// 	var bookings []string // <= a slice

// 	/*  -------- map ---------
// 	*	map is a collection of key value pairs
// 	*	[
// 	*		"full_name": "ppm",
// 	*	]
// 	*	all keys have the same data type
// 	*	all values have the same data type
// 	*	can't mix data type for value
// 	*	to create an empty map => make(map[string]string)
// 	*	creating a slice of maps => var users = make([]map[string]string, size) <= need to assign initial size of the slice
// 	 */

// 	var user_data = make(map[string]string)
// 	user_data["name"] = "PPM"
// 	user_data["email"] = "ppm@m.co"
// 	// user_data["age"] = 27 // this will give an error because of different data type
// 	user_data["age"] = strconv.FormatUint(uint64(27), 10) // convert uint data to string

// 	var users = make([]map[string]string, 0) // <= this is a slice of maps

// 	users = append(users, user_data)

// 	fmt.Printf("maps %v \n", users)

// 	fmt.Printf("conference_name is a %T, conference_tickets is an %T, remaining_tickets is an %T\n", conference_name, conference_tickets, remaining_tickets)

// 	fmt.Println("Get tickets")

// 	// for { // for without any thing is an infinite loop
// 	for remaining_tickets > 0 {

// 		var full_name string
// 		var email string
// 		var user_tickets uint

// 		fmt.Println(&remaining_tickets) // &remainingTickets is a pointer which is a special variable to point the memory location of a variable

// 		// when using scan, asking for user input, need to use pointer "&var"
// 		fmt.Println("Enter Full Name:")
// 		fmt.Scan(&full_name)
// 		fmt.Println("Enter Email:")
// 		fmt.Scan(&email)
// 		fmt.Println("Enter Number of Tickets:")
// 		fmt.Scan(&user_tickets)

// 		// is_valid_name, is_valid_email, is_valid_ticket := validateInputs(full_name, email, remaining_tickets, user_tickets) // moved to helper file

// 		is_valid_name, is_valid_email, is_valid_ticket := validator.ValidateInputsFromValidator(full_name, email, remaining_tickets, user_tickets) // when importing from another package, need to use like this => packagename.functionname()

// 		// if remaining_tickets < user_tickets {
// 		// 	fmt.Printf("we only have %v, can't book %v tickets", remaining_tickets, user_tickets)
// 		// 	continue
// 		// 	// break will stop the loop
// 		// 	// continue will skip all the code below and restart the next iteration
// 		// }

// 		if is_valid_name && is_valid_email && is_valid_ticket {
// 			// invalid operation will if the variables in the calculation do not have the same types
// 			remaining_tickets = remaining_tickets - user_tickets
// 			// bookings[0] = full_name // use this way to add data to array
// 			bookings = append(bookings, full_name) // use this way to add data to slice
// 			// bookings[52] = full_name // out of bond error will show

// 			fmt.Printf("Array: %v\n", bookings)
// 			fmt.Printf("First Element: %v\n", bookings[0])
// 			fmt.Printf("Array Type: %T\n", bookings)
// 			fmt.Printf("Array length: %v\n", len(bookings))
// 			fmt.Printf("%v booked %v tickets. Confirmation mail will be sent to %v\n", full_name, user_tickets, email)
// 			fmt.Println("Remaining Tickets: ", remaining_tickets)

// 			printFirstNames(bookings)

// 			// var no_tickets bool = remaining_tickets == 0 // boolean data type
// 			if remaining_tickets == 0 {
// 				fmt.Println("Booked out!")
// 				break
// 			}

// 		} else {
// 			fmt.Println("Invalid Entry! Try again.")
// 			continue
// 		}

// 	}

// 	// city := "Bangkok"
// 	// switch city {
// 	// 	case "Singapore", "Bangkok":
// 	// 		fmt.Printf("want to work in %v", city)
// 	// 	case "Lashio", "Hopin":
// 	// 		fmt.Printf("hometown")
// 	// 	case "Mandalay":
// 	// 		fmt.Printf("live in %v", city)
// 	// 	default:
// 	// 		fmt.Printf("unkonwn")
// 	// }

// }

// func greetUser() {
// 	fmt.Printf("Welcome to %v dude! \n", conference_name)
// 	fmt.Printf("Total Seats: %v\nRemaining Tickets: %v\n", conference_tickets, remaining_tickets)
// }

// /*
// *	watch the tuto below for pointer and how it's useful for function
// *	https://www.youtube.com/watch?v=sTFJtxJXkaY
//  */

// func printFirstNames(bookings []string) {
// 	first_names := []string{}
// 	// range iterates for different data structure, not only for iterating array and slices
// 	// for array & slices, range provides index and value for each element
// 	// _ is called blank identifier, used for unused variable and can also be used to call only init function of another package
// 	for _, ele := range bookings {
// 		var names = strings.Fields(ele) // strings.Fields() splits the string at ' ', strings.Fields('Pyae Phyo M g') => ['Pyae', 'Phyo', 'M', 'g']
// 		first_names = append(first_names, names[0])
// 	}
// 	fmt.Printf("First Names: %v", first_names)
// }

// /*
// *	function in GO
// *	func functionName(param1 param1_data_type, param2 param2_data_type,... ) (return_data_type,...)
// *
// *	fucntion with one return value
// *	func validateInputs(full_name string, email string, user_tickets string) bool {
// *		return true
// *	}
// *
// *	unlike other languages GO functions can return multiple values
// *
// *	fucntion with one return value
// *	func validateInputs(full_name string, email string, user_tickets string) (bool, bool, bool) {
// *		return true, true, true
// *	}
//  */
