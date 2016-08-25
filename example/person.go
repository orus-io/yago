package main

import (
	"time"

	"github.com/orus-io/yago"
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
	PersonID int64 `yago:"fk=Person"`
	Name     string
	Number   string
}

// BeforeInsert callback
func (p *Person) BeforeInsert(db *yago.DB) {
	p.CreatedAt = time.Now()
}

// BeforeUpdate callback
func (p *Person) BeforeUpdate(db *yago.DB) {
	now := time.Now()
	p.UpdatedAt = &now
}

// NewPerson instantiate a Person with sensible default values
func NewPerson() *Person {
	return &Person{}
}
