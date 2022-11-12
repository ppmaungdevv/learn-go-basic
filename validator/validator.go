package validator

import "strings"

// when you want to export varible or function, declare the first letter with capital letter => Exportvar, Exportfunc

func ValidateInputsFromValidator(full_name string, email string, remaining_tickets uint, user_tickets uint) (bool, bool, bool) {
	is_valid_name := len(full_name) >= 2
	is_valid_email := strings.Contains(email, "@")
	is_valid_ticket := user_tickets > 0 && remaining_tickets < user_tickets
	return is_valid_name, is_valid_email, is_valid_ticket
}
