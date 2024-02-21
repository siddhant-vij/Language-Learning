package utils

import (
	"regexp"
	"strconv"

	"go-learning-14/pkg/contact"
)

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func ValidatePhone(phone int) bool {
	phoneRegex := regexp.MustCompile(`^\d{10}$`)
	return phoneRegex.MatchString(strconv.Itoa(phone))
}

func ValidateContact(c contact.Contact) bool {
	return ValidateEmail(c.Email) && ValidatePhone(c.PhoneNumber)
}
