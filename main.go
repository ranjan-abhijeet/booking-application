package main

import (
	"fmt"
	"strings"
)

func main() {
	var conference_name string = "Go Conference"
	const total_tickets uint16 = 100
	var remaining_tickets uint16 = 100
	var bookings []string // Array or Slice of string type. bookings := []string{}

	welcome(conference_name, total_tickets, remaining_tickets)

	for {
		var no_tickets bool = remaining_tickets == 0
		if no_tickets {
			fmt.Printf("Our %v is booked out.", conference_name)
			fmt.Println("Come back next year.")
			break
		}
		first_name, last_name, email, _, ticket_numbers := get_user_input()

		is_valid_name, is_valid_email, is_valid_ticket_number := validate_user_input(first_name, last_name, email, ticket_numbers, remaining_tickets)

		if is_valid_ticket_number && is_valid_name && is_valid_email {
			book_tickets(ticket_numbers, total_tickets, remaining_tickets, bookings, first_name, last_name, email)
			var first_names = get_attendees(bookings)
			fmt.Printf("Participant's list: %v\n", first_names)

		} else {

			if !is_valid_name {
				fmt.Printf("Your entered name: %v %v is invalid.\n", first_name, last_name)
			}
			if !is_valid_email {
				fmt.Printf("Your entered email: %v is invalid.\n", email)
			}
			if !is_valid_ticket_number {
				fmt.Printf("Your entered number of ticket: %v is invalid.\n", ticket_numbers)
			} else {
				fmt.Println("Sorry, could not process your response. Try again!")
				continue
			}
		}
	}
}

func welcome(conference string, total_tickets uint16, remaining_tickets uint16) {
	fmt.Printf("Welcome to the %v booking application\n", conference)
	fmt.Printf("Total tickets: %v\nAvailable tickets: %v\n", total_tickets, remaining_tickets)
}

func get_attendees(booking_list []string) []string {
	var first_names []string
	for _, booking := range booking_list {
		var name = strings.Fields(booking)
		first_names = append(first_names, name[0])
	}
	return first_names
}

func validate_user_input(first_name string, last_name string, email string, ticket_numbers uint16, remaining_tickets uint16) (bool, bool, bool) {
	var is_valid_name bool = len(first_name) >= 2 && len(last_name) >= 2
	var is_valid_email bool = strings.Contains(email, "@")
	var is_valid_ticket_number bool = ticket_numbers > 0 && ticket_numbers <= remaining_tickets
	return is_valid_name, is_valid_email, is_valid_ticket_number
}

func get_user_input() (string, string, string, string, uint16) {
	var first_name string
	var last_name string
	var email string
	var ticket_numbers uint16
	var conference_city string

	fmt.Println()
	fmt.Print("Enter your first name: ")
	fmt.Scan(&first_name)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&last_name)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter the city where you want to attend the conference: ")
	fmt.Scan(&conference_city)

	switch conference_city {
	case "New York":
		fmt.Printf("Welcome to %v\n", conference_city)
	case "Singapore":
		fmt.Printf("Welcome to %v\n", conference_city)
	case "London":
		fmt.Printf("Welcome to %v\n", conference_city)
	case "Mumbai":
		fmt.Printf("Welcome to %v\n", conference_city)
	case "Chandigarh":
		fmt.Printf("Welcome to %v\n", conference_city)
	default:
		fmt.Printf("%v is not a valid location for the conference\n", conference_city)
		fmt.Println("Try again.")
		return get_user_input()
	}

	fmt.Print("Enter the number of tickets you need: ")
	fmt.Scan(&ticket_numbers)

	return first_name, last_name, email, conference_city, ticket_numbers
}

func book_tickets(ticket_numbers uint16, total_tickets uint16, remaining_tickets uint16, bookings []string, first_name string, last_name string, email string) {
	fmt.Printf("\nThank you %v %v for booking %v tickets.\nYou will receive a confirmation email on %v\n", first_name, last_name, ticket_numbers, email)

	remaining_tickets = remaining_tickets - ticket_numbers

	fmt.Printf("Total tickets: %v\nAvailable tickets: %v\n", total_tickets, remaining_tickets)
	bookings = append(bookings, first_name+" "+last_name)
}
