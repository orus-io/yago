package main

import (
	"reflect"
	"time"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yagorm"
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

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (Person) StructType() reflect.Type {
	return personType
}

// PersonMapper is the Person mapper
type PersonMapper struct{}

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

// Values returns the struct values as a map
func (mapper PersonMapper) Values(instance yagorm.MappedStruct) map[string]interface{} {
	s := instance.(*Person)
	if s == nil {
		panic("Wrong struct type passed to the mapper")
	}
	return map[string]interface{}{
		"id":         s.ID,
		"name":       s.Name,
		"email":      s.Email,
		"created_at": s.CreatedAt,
		"updated_at": s.UpdatedAt,
	}
}
