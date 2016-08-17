package main

import (
	"reflect"
	"time"

	"github.com/aacanakin/qb"
)

var personTable = qb.Table(
	"person",
	qb.Column("id", qb.Varchar().Size(36)),
	qb.Column("name", qb.Varchar().NotNull()),
	qb.Column("email", qb.Varchar()),
	qb.Column("created_at", qb.Timestamp().NotNull()),
	qb.Column("updated_at", qb.Timestamp()),
	qb.PrimaryKey("id"),
)

var personType = reflect.TypeOf(Person{})

// NewPerson instanciate a Person with sensible default values
func NewPerson() *Person {
	return &Person{
		CreatedAt: time.Now(),
	}
}

// Values returns the struct values as a map
func (p Person) Values() map[string]interface{} {
	return map[string]interface{}{
		"id":         p.ID,
		"name":       p.Name,
		"email":      p.Email,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}
}

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (*Person) StructType() reflect.Type {
	return personType
}

// PersonMapper is the Person mapper
type PersonMapper struct {
}

// Name returns the mapper name
func (*PersonMapper) Name() string {
	return "example/person"
}

// Table returns the mapper table
func (*PersonMapper) Table() *qb.TableElem {
	return &personTable
}

// StructType returns the reflect.Type of the mapped structure
func (PersonMapper) StructType() reflect.Type {
	return personType
}
