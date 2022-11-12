package main

/*
*
*	in GO, code can be split into multiple files under the same package name
*	when running multiple GO file, same package,
*	1. use "go run ." otherwise all the file names must be specified like "go run main.go helper.go"
*	2. functions inside separate file with the same package does not need to use with import
 */

import "strings"

func validateInputs(full_name string, email string, remaining_tickets uint, user_tickets uint) (bool, bool, bool) {
	is_valid_name := len(full_name) >= 2
	is_valid_email := strings.Contains(email, "@")
	is_valid_ticket := user_tickets > 0 && remaining_tickets < user_tickets
	return is_valid_name, is_valid_email, is_valid_ticket
}
