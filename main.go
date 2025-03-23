package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	authorName := "Richard"
	bookings := []string{}

	fmt.Println(remainingTickets)
	fmt.Println(&remainingTickets)

	// Print variables types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking system\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Please select an option")
	fmt.Printf("The author name is : %v\n", authorName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their first name
		fmt.Println("Enter your first name :")
		fmt.Scan(&firstName)

		// ask user for their last name
		fmt.Println("Enter your last name :")
		fmt.Scan(&lastName)

		// ask user for their email

		fmt.Println("Enter your email :")
		fmt.Scan(&email)

		// ask user for the number of tickets they want to buy
		fmt.Println("Enter number of tickets you want to buy :")
		fmt.Scan(&userTickets)
		isValideName := len(firstName) >= 2 && len(lastName) >= 2
		isValideEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValideTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		moreTickets := userTickets > remainingTickets
		if isValideName && isValideEmail && isValideTicketsNumber {
			fmt.Println("Sorry, we don't have enough tickets for you")
			continue
		}

		fmt.Printf("More ticket %v", moreTickets)
		// calcul remaining tickets
		remainingTickets = remainingTickets - userTickets
		bookings[0] = firstName + " " + lastName

		// print booking details
		fmt.Printf("The whole array : %v\n", bookings)
		fmt.Printf("The first booking : %v\n", bookings[0])
		fmt.Printf("Array type : %T\n", bookings)
		fmt.Printf("Array length : %v\n", len(bookings))

		// thanks message
		fmt.Printf("Thank you %v %v, for booking. You will receive a confirmation email at %v\n", firstName, lastName, email)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are : %v\n", firstNames)
		// print remaining tickets
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		var noTicketsRemaining bool = remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("No more tickets available")
			break
		}
	}
}
