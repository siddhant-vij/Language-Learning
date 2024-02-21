package store

import (
	"encoding/json"
	"errors"
	"os"

	"go-learning-14/pkg/contact"
	"go-learning-14/utils"
)

type Store struct {
	Contacts map[int]contact.Contact
}

func NewStore() *Store {
	return &Store{
		Contacts: make(map[int]contact.Contact),
	}
}

func (s *Store) AddContact(c contact.Contact) error {
	if !utils.ValidateContact(c) {
		return errors.New("invalid contact details")
	}

	if _, exists := s.Contacts[c.PhoneNumber]; exists {
		return errors.New("contact already exists")
	}

	s.Contacts[c.PhoneNumber] = c
	return s.SaveToFile()
}

func (s *Store) GetContact(phoneNumber int) (contact.Contact, bool) {
	contact, exists := s.Contacts[phoneNumber]
	return contact, exists
}

func (s *Store) UpdateContact(phoneNumber int, updatedContact contact.Contact) error {
	if _, exists := s.Contacts[phoneNumber]; !exists {
		return errors.New("contact does not exist")
	}

	if !utils.ValidateContact(updatedContact) {
		return errors.New("invalid contact details")
	}

	s.Contacts[phoneNumber] = updatedContact
	return s.SaveToFile()
}

func (s *Store) DeleteContact(phoneNumber int) error {
	if _, exists := s.Contacts[phoneNumber]; !exists {
		return errors.New("contact does not exist")
	}

	delete(s.Contacts, phoneNumber)
	return s.SaveToFile()
}

func (s *Store) ListContacts() []contact.Contact {
	contactsList := make([]contact.Contact, 0, len(s.Contacts))
	for _, c := range s.Contacts {
		contactsList = append(contactsList, c)
	}
	return contactsList
}

func (s *Store) SearchContacts(searchTerm int) []contact.Contact {
	searchResults := []contact.Contact{}
	for _, c := range s.Contacts {
		if utils.ValidatePhone(searchTerm) && c.PhoneNumber == searchTerm {
			searchResults = append(searchResults, c)
		}
	}
	return searchResults
}

func (s *Store) LoadFromFile() error {
	file, err := os.ReadFile("data/contacts.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &s.Contacts)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) SaveToFile() error {
	file, err := json.MarshalIndent(s.Contacts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data/contacts.json", file, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
