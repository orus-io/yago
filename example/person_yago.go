package main

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
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

// PersonModel
type PersonModel struct {
	mapper *PersonMapper
	ID yago.ScalarField
	Name yago.ScalarField
	Email yago.ScalarField
	CreatedAt yago.ScalarField
	UpdatedAt yago.ScalarField
}

func NewPersonModel(meta *yago.Metadata) PersonModel {
	mapper := NewPersonMapper()
	meta.AddMapper(mapper)
	return PersonModel {
		mapper: mapper,
		ID: yago.NewScalarField(mapper.Table().C("id")),
		Name: yago.NewScalarField(mapper.Table().C("name")),
		Email: yago.NewScalarField(mapper.Table().C("email_address")),
		CreatedAt: yago.NewScalarField(mapper.Table().C("created_at")),
		UpdatedAt: yago.NewScalarField(mapper.Table().C("updated_at")),
	}
}

func (m PersonModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewPersonMapper initialize a NewPersonMapper
func NewPersonMapper() *PersonMapper {
	m := &PersonMapper{}
	return m
}

// PersonMapper is the Person mapper
type PersonMapper struct{}

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

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper PersonMapper) SQLValues(instance yago.MappedStruct) map[string]interface{} {
	s, ok := instance.(*Person)
	if !ok {
		 panic("Wrong struct type passed to the mapper")
	}
	m := make(map[string]interface{})
	if s.ID != 0 {
		m["id"] = s.ID
	}
	m["name"] = s.Name
	m["email_address"] = s.Email
	m["created_at"] = s.CreatedAt
	m["updated_at"] = s.UpdatedAt
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


var phoneNumberTable = qb.Table(
	"phonenumber",
	qb.Column("id", qb.BigInt()).PrimaryKey().AutoIncrement(),
	qb.Column("person_id", qb.BigInt()),
	qb.Column("name", qb.Varchar().NotNull()),
	qb.Column("number", qb.Varchar().NotNull()),
	qb.ForeignKey().Ref("person_id", "person", "id"),
)

var phoneNumberType = reflect.TypeOf(PhoneNumber{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (PhoneNumber) StructType() reflect.Type {
	return phoneNumberType
}

// PhoneNumberModel
type PhoneNumberModel struct {
	mapper *PhoneNumberMapper
	ID yago.ScalarField
	PersonID yago.ScalarField
	Name yago.ScalarField
	Number yago.ScalarField
}

func NewPhoneNumberModel(meta *yago.Metadata) PhoneNumberModel {
	mapper := NewPhoneNumberMapper()
	meta.AddMapper(mapper)
	return PhoneNumberModel {
		mapper: mapper,
		ID: yago.NewScalarField(mapper.Table().C("id")),
		PersonID: yago.NewScalarField(mapper.Table().C("person_id")),
		Name: yago.NewScalarField(mapper.Table().C("name")),
		Number: yago.NewScalarField(mapper.Table().C("number")),
	}
}

func (m PhoneNumberModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewPhoneNumberMapper initialize a NewPhoneNumberMapper
func NewPhoneNumberMapper() *PhoneNumberMapper {
	m := &PhoneNumberMapper{}
	return m
}

// PhoneNumberMapper is the PhoneNumber mapper
type PhoneNumberMapper struct{}

// Name returns the mapper name
func (*PhoneNumberMapper) Name() string {
	return "main/PhoneNumber"
}

// Table returns the mapper table
func (*PhoneNumberMapper) Table() *qb.TableElem {
	return &phoneNumberTable
}

// StructType returns the reflect.Type of the mapped structure
func (PhoneNumberMapper) StructType() reflect.Type {
	return phoneNumberType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper PhoneNumberMapper) SQLValues(instance yago.MappedStruct) map[string]interface{} {
	s, ok := instance.(*PhoneNumber)
	if !ok {
		 panic("Wrong struct type passed to the mapper")
	}
	m := make(map[string]interface{})
	if s.ID != 0 {
		m["id"] = s.ID
	}
	m["person_id"] = s.PersonID
	m["name"] = s.Name
	m["number"] = s.Number
	return m
}

// FieldList returns the list of fields for a select
func (mapper PhoneNumberMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		phoneNumberTable.C("id"),
		phoneNumberTable.C("person_id"),
		phoneNumberTable.C("name"),
		phoneNumberTable.C("number"),
	}
}

// Scan a struct
func (mapper PhoneNumberMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*PhoneNumber)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.ID,
		&s.PersonID,
		&s.Name,
		&s.Number,
	)
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PhoneNumberMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return phoneNumberTable.C("id").Eq(instance.(*PhoneNumber).ID)
}
