package contact

import (
	"fmt"
)

type Contact struct {
	Name        string
	Email       string
	PhoneNumber int
}

func (c *Contact) Format() string {
	return fmt.Sprintf("Name: %s\nEmail: %s\nPhone: %d", c.Name, c.Email, c.PhoneNumber)
}
