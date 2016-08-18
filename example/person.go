package main

import (
	"time"
)

//go:generate yagorm

// Person is a person record in the database
//yagorm:
type Person struct {
	ID        int64 `yagorm:"primary_key"`
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
