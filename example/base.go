package main

//go:generate yago

// Base struct for all model structs
//yago:notable
type Base struct {
	ID int64 `yago:"primary_key,auto_increment"`
}
