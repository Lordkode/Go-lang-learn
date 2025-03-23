package helper

import "strings"

func validateUserInput(firsName string, lastName string, email string, userTickets uint) {
	isValideName := len(firsName) >= 2 && len(lastName) >= 2
	isValideEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValideTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValideName && isValideEmail && isValideTicketsNumber
}
