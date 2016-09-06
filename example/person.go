package main

import ()

//go:generate yago

// Person is a person record in the database
//yago:autoattrs
type Person struct {
	Base
	Name  string  `yago:"index"`
	Email *string `yago:"email_address,unique_index"`
}

// PhoneNumber is a phone number
//yago:autoattrs
type PhoneNumber struct {
	Base
	PersonID int64 `yago:"fk=Person ondelete cascade onupdate cascade"`
	Name     string
	Number   string
}

// NewPerson instantiate a Person with sensible default values
func NewPerson() *Person {
	return &Person{}
}
