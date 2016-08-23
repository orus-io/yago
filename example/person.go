package main

import (
	"time"
)

//go:generate yago

// Person is a person record in the database
//yago:autoattrs
type Person struct {
	ID        int64   `yago:"primary_key,auto_increment"`
	Name      string  `yago:"index"`
	Email     *string `yago:"email_address,unique_index"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// PhoneNumber is a phone number
//yago:autoattrs
type PhoneNumber struct {
	ID       int64 `yago:"primary_key,auto_increment"`
	PersonID int64 //`yago:"fk=Person.ID"`
	Name     string
	Number   string
}

// NewPerson instantiate a Person with sensible default values
func NewPerson() *Person {
	return &Person{
		CreatedAt: time.Now(),
	}
}
