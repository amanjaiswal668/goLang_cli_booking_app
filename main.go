package main

import (
	"cli_booking_app/helper"
	"fmt"
	"sync"
	"time"
)

// Packages level variables. It can  be accessed by ny function in a package.
var conferenceName = "Go lang Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

// var bookings = []string{}
// List of Dictionary
// var bookings = make([]map[string]string, 0)
// Using Struct
var bookings = make([]UserDataStruct, 0)

// Custom  type of Datatypes.
// Struct, (What user type would like, pre-defined)
// Struct is what classes in JAVA.
type UserDataStruct struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Create a wait group for making the main thread to wait for the complete execution of the child thread if any.
var waitGroup = sync.WaitGroup{}

func main() {
	fmt.Print("Hello User!\n")

	// Function level variables, can only be acces by parent function and can be passed to a child function.
	// Variable declearation in go. Type 2 can only be used for var not for const.
	// 1. var conferenceName = "Go lang Conference"
	// 2. conferenceName := "Go lang Conference"

	// var conferenceName = "Go lang Conference"
	// const conferenceTickets uint = 50
	// var remainingTickets uint = 50

	// // var bookings [50]string -> For Array.
	// var bookings []string // -> For slices (Dynamic sized array.) List of userNames in string

	// Print values in console.
	// 1. Using println(), it uses comma and variable name for displaying the variables.
	// 2. Using printf(), it uses %v for assigning the variables.

	fmt.Printf("Types of the varialbes :- \n")
	fmt.Printf("ConferenceName is %T, conferenceTicket is %T\n\n", conferenceName, conferenceTickets)

	// greetUser(conferenceName, conferenceTickets, remainingTickets)
	greetUser()
	// Replaced with function greetUser
	// fmt.Println("Welcome to ", conferenceName, ", a CLI based booking applications.")
	// fmt.Printf("Welcome to %v, a CLI based booking applications.\n", conferenceName)
	// fmt.Println("We have total of ", conferenceTickets, " and ", remainingTickets, "are remaining.")

	// Loops. -> Control flow.
	// for loop
	for {
		// replaced with the function getUserInput
		// Asking user inputs inside the for loop. Basically replication of a real time booking app where it does not end.
		// var firstName string
		// var lastName string
		// var email string
		// var userTickets uint
		// // Ask for the user name.
		// fmt.Printf("Enter the first name :-  \n")
		// fmt.Scan(&firstName)
		// fmt.Printf("Enter the last name :-  \n")
		// fmt.Scan(&lastName)
		// fmt.Printf("Enter the email name :-  \n")
		// fmt.Scan(&email)
		// fmt.Printf("Enter the number of tickets :-  \n")
		// fmt.Scan(&userTickets)

		firstName, lastName, email, userTickets := getUserInput()

		// Validate the user input.
		// Replaced with function validateUserInput
		// isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email, "@")
		// isValidTicketNumber := userTickets > 0
		// var isValidTicketNumber := userTickets > 0 && userTickets > remainingTickets

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		// if isValidName && isValidEmail && isValidTicketNumber && userTickets > remainingTickets {
		// 	fmt.Printf("We have only %v tickets available. Please try again. \n", remainingTickets)
		// 	// Throw user out of the app.
		// 	// break
		// 	// Ask user for fresh input.
		// 	continue
		// }
		if isValidName && isValidEmail && isValidTicketNumber {

			// remainingTickets = remainingTickets - userTickets
			// // bookings[0] = firstName + " " + lastName
			// // Slices is something similar to ArrayList in JAVA. It is a dynamic sized array, it internally uses array.
			// bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("User %v %v has been created. \n", firstName, lastName)
			// fmt.Printf("User %v %v has in total %v tickets booked.\n", firstName, lastName, userTickets)

			bookTicket(firstName, lastName, email, userTickets)
			// Single threaded.
			// sendTicket(userTickets, firstName, lastName, email)
			// Multi threaded.(Go-routine)
			// Will start multiple threads based on the number of calls.
			waitGroup.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			// Replaced with function printFirstNames
			// firstNames := []string{}
			// for _, booking := range bookings {
			// 	var names = strings.Fields(booking)
			// 	firstNames = append(firstNames, names[0])
			// }
			// fmt.Printf("Our bookings are :- \n %v \n", firstNames)

			// Call function to print firstNames.

			firstNames := printFirstNames()
			fmt.Printf("Our bookings are :- \n %v \n", firstNames)

			// Booleans. Declearation of Booleans.
			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			// if noTicketsRemaining == 0 {

			if remainingTickets == 0 {
				// End the program.
				fmt.Printf("Booking closed!")
				break
			}

			fmt.Printf("%v tickets remaining \n", remainingTickets)
		}
		// else {
		// 	is !isValidName {

		// 	}
		// }

	}
	waitGroup.Wait()
}

// Greet user.
// Accessing package level variables
func greetUser() {
	fmt.Println("Welcome to ", conferenceName, ", a CLI based booking applications.")
	fmt.Printf("Welcome to %v, a CLI based booking applications.\n", conferenceName)
	fmt.Println("We have total of ", conferenceTickets, " and ", remainingTickets, "are remaining.")
}

// // Greet user.
// Accessing function level variables
// func greetUser(conferenceName string, conferenceTickets uint, remainingTickets uint) {
// 	fmt.Println("Welcome to ", conferenceName, ", a CLI based booking applications.")
// 	fmt.Printf("Welcome to %v, a CLI based booking applications.\n", conferenceName)
// 	fmt.Println("We have total of ", conferenceTickets, " and ", remainingTickets, "are remaining.")
// }

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firsName"]) // Accessing firstName from the map.
		firstNames = append(firstNames, booking.firstName) // Accessing firstName from the struct.
	}
	// fmt.Printf("Our bookings are :- \n %v \n", firstNames)
	return firstNames
}

// Get user Input

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask for the user name.
	fmt.Printf("Enter the first name :-  \n")
	fmt.Scan(&firstName)
	fmt.Printf("Enter the last name :-  \n")
	fmt.Scan(&lastName)
	fmt.Printf("Enter the email name :-  \n")
	fmt.Scan(&email)
	fmt.Printf("Enter the number of tickets :-  \n")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Book Tickets.

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// // Create a map{Dictionary in general} for a user. Can only have similar dataType data.
	// var userData = make(map[string]string)
	// // Add all user data to userData.
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// 1. strconv -> A builtin method to convert other datatype to string.
	// 2. FormatUnit -> Formats the input integer.
	// 3. unit64 -> Takes the integer value and convert it into its corresponding type.
	//  It takes 2 values as a parameter, first is integer value and second is base value.

	// Using UserDataStruct

	var userData = UserDataStruct{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// bookings = append(bookings, firstName+" "+lastName)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings :- \n %v \n", bookings)

	fmt.Printf("User %v %v has been created. \n", firstName, lastName)
	fmt.Printf("User %v %v has in total %v tickets booked.\n", firstName, lastName, userTickets)

	// return bookings
}

// Go-routine(thread)
// Simulating the long process.
// Send ticket to the user.

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("*************************************************** \n")
	fmt.Printf("Sending ticket: \n %v \n to email %v \n", ticket, email)
	fmt.Printf("*************************************************** \n")
	waitGroup.Done()
}
