package main

import (
	"time"
)

//go:generate yagorm

// Person is a person record in the database
//yagorm:autoattrs
type Person struct {
	ID        int64 `yagorm:"primary_key,auto_increment"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// NewPerson instanciate a Person with sensible default values
func NewPerson() *Person {
	return &Person{
		CreatedAt: (time.Now()),
	}
}
