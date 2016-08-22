package main

import (
	"time"
)

//go:generate yago

// Person is a person record in the database
//yago:autoattrs
type Person struct {
	ID        int64 `yago:"primary_key,auto_increment"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// NewPerson instantiate a Person with sensible default values
func NewPerson() *Person {
	return &Person{
		CreatedAt: time.Now(),
	}
}
