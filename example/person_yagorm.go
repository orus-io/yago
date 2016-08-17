package main

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yagorm"
)

var personTable = qb.Table(
	"person",
	qb.Column("id", qb.Int().AutoIncrement()),
	qb.Column("name", qb.Varchar().NotNull()),
	qb.Column("email", qb.Varchar()),
	qb.Column("created_at", qb.Timestamp().NotNull()),
	qb.Column("updated_at", qb.Timestamp()),
	qb.PrimaryKey("id"),
)

var personType = reflect.TypeOf(Person{})

// Dirty workaround for qb not supporting auto increment on sqlite.
// VERY temporary
var personLastID int64

// NewPerson instanciate a Person with sensible default values
func NewPerson() *Person {
	personLastID++
	return &Person{
		ID:        personLastID,
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

// Values returns non-default values as a map
func (mapper PersonMapper) Values(instance yagorm.MappedStruct) map[string]interface{} {
	s := instance.(*Person)
	if s == nil {
		panic("Wrong struct type passed to the mapper")
	}
	m := make(map[string]interface{})
	if s.ID != 0 {
		m["id"] = s.ID
	}
	if s.Name != "" {
		m["name"] = s.Name
	}
	if s.Email != nil {
		m["email"] = s.Email
	}
	if s.CreatedAt != (time.Time{}) {
		m["created_at"] = s.CreatedAt
	}
	if s.UpdatedAt != nil {
		m["updated_at"] = s.UpdatedAt
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personTable.C("id"),
		personTable.C("name"),
		personTable.C("email"),
		personTable.C("created_at"),
		personTable.C("updated_at"),
	}
}

// Scan a struct
func (mapper PersonMapper) Scan(rows *sql.Rows, instance yagorm.MappedStruct) error {
	s := instance.(*Person)
	if s == nil {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(&s.ID, &s.Name, &s.Email, &s.CreatedAt, &s.UpdatedAt)
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonMapper) PKeyClause(instance yagorm.MappedStruct) qb.Clause {
	return personTable.C("id").Eq(instance.(*Person).ID)
}
