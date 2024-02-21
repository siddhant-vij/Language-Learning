package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-learning-14/pkg/contact"
	"go-learning-14/pkg/store"
	"go-learning-14/utils"
)

func main() {
	// Initialize store
	contactStore := store.NewStore()
	err := contactStore.LoadFromFile()
	if err != nil {
		fmt.Printf("Error loading contacts: %s\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("\n\nContact Manager CLI")
		fmt.Printf("\n-------------------\n\n")
		fmt.Println("Choose an option:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Delete Contact")
		fmt.Println("3. Update Contact")
		fmt.Println("4. List Contacts")
		fmt.Println("5. Search Contact")
		fmt.Println("6. Exit")
		fmt.Print("Enter option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			newContact, exitStatus := readContactInput()
			if exitStatus {
				continue
			}
			err := contactStore.AddContact(newContact)
			handleError(err)
		case "2":
			fmt.Print("Enter contact phone number to delete: ")
			scanner.Scan()
			phoneNumber, err := strconv.Atoi(scanner.Text())
			handleError(err)
			err = contactStore.DeleteContact(phoneNumber)
			handleError(err)
		case "3":
			fmt.Print("Enter contact phone number to update: ")
			scanner.Scan()
			phoneNumber, err := strconv.Atoi(scanner.Text())
			handleError(err)
			updatedContact, exitStatus := readContactInput()
			if exitStatus {
				continue
			}
			err = contactStore.UpdateContact(phoneNumber, updatedContact)
			handleError(err)
		case "4":
			contacts := contactStore.ListContacts()
			fmt.Println("Contacts List:")
			for _, c := range contacts {
				fmt.Printf("%v\n", c)
			}
		case "5":
			fmt.Print("Enter search term: ")
			scanner.Scan()
			searchTerm, err := strconv.Atoi(scanner.Text())
			handleError(err)
			results := contactStore.SearchContacts(searchTerm)
			fmt.Println("Search Results:")
			for _, c := range results {
				fmt.Printf("%v\n", c)
			}
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func readContactInput() (contact.Contact, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	var c contact.Contact
	var err error

	for {
		fmt.Print("Enter contact name (or type 'exit' to return to main menu): ")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			return contact.Contact{}, true
		}
		c.Name = input

		fmt.Print("Enter contact phone number (or type 'exit' to return to main menu): ")
		scanner.Scan()
		input = scanner.Text()
		if strings.ToLower(input) == "exit" {
			return contact.Contact{}, true
		}
		c.PhoneNumber, err = strconv.Atoi(input)
		if err != nil {
			handleError(err)
			continue
		}
		if !utils.ValidatePhone(c.PhoneNumber) {
			fmt.Println("Invalid phone number. Please try again.")
			continue
		}

		fmt.Print("Enter contact email (or type 'exit' to return to main menu): ")
		scanner.Scan()
		input = scanner.Text()
		if strings.ToLower(input) == "exit" {
			return contact.Contact{}, true
		}
		c.Email = input
		if !utils.ValidateEmail(c.Email) {
			fmt.Println("Invalid email. Please try again.")
			continue
		}

		break
	}

	return c, false
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
}
