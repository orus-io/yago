package main

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/aacanakin/qb"

	"bitbucket.org/cdevienne/yago"
)



var personTable = qb.Table(
	"person",
	qb.Column("id", qb.BigInt()).PrimaryKey().AutoIncrement(),
	qb.Column("name", qb.Varchar().NotNull()),
	qb.Column("email_address", qb.Varchar()),
	qb.Column("created_at", qb.Timestamp().NotNull()),
	qb.Column("updated_at", qb.Timestamp()),
	qb.UniqueKey(
		"email_address",
	),
).Index(
	"name",
)

var personType = reflect.TypeOf(Person{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (Person) StructType() reflect.Type {
	return personType
}

// PersonFields
type PersonFields struct {
	ID qb.ColumnElem
	Name qb.ColumnElem
	Email qb.ColumnElem
	CreatedAt qb.ColumnElem
	UpdatedAt qb.ColumnElem
}

// NewPersonMapper initialize a NewPersonMapper
func NewPersonMapper() *PersonMapper {
	m := &PersonMapper{}
	m.Fields.ID = m.Table().C("id")
	m.Fields.Name = m.Table().C("name")
	m.Fields.Email = m.Table().C("email_address")
	m.Fields.CreatedAt = m.Table().C("created_at")
	m.Fields.UpdatedAt = m.Table().C("updated_at")
	return m
}

// PersonMapper is the Person mapper
type PersonMapper struct{
	Fields PersonFields
}

// Name returns the mapper name
func (*PersonMapper) Name() string {
	return "main/Person"
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
func (mapper PersonMapper) Values(instance yago.MappedStruct) map[string]interface{} {
	s, ok := instance.(*Person)
	if !ok {
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
		m["email_address"] = s.Email
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
		personTable.C("email_address"),
		personTable.C("created_at"),
		personTable.C("updated_at"),
	}
}

// Scan a struct
func (mapper PersonMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*Person)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.ID,
		&s.Name,
		&s.Email,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return personTable.C("id").Eq(instance.(*Person).ID)
}
