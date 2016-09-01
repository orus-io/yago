package main

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
)

const (
	PersonTableName = "person"
	PersonIDColumnName = "id"
	PersonNameColumnName = "name"
	PersonEmailColumnName = "email_address"
	PersonCreatedAtColumnName = "created_at"
	PersonUpdatedAtColumnName = "updated_at"
)

var personTable = qb.Table(
	PersonTableName,
	qb.Column(PersonIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(PersonNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PersonEmailColumnName, qb.Varchar()).Null(),
	qb.Column(PersonCreatedAtColumnName, qb.Timestamp()).NotNull(),
	qb.Column(PersonUpdatedAtColumnName, qb.Timestamp()).Null(),
	qb.UniqueKey(
		PersonEmailColumnName,
	),
).Index(
	PersonNameColumnName,
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
		ID: yago.NewScalarField(mapper.Table().C(PersonIDColumnName)),
		Name: yago.NewScalarField(mapper.Table().C(PersonNameColumnName)),
		Email: yago.NewScalarField(mapper.Table().C(PersonEmailColumnName)),
		CreatedAt: yago.NewScalarField(mapper.Table().C(PersonCreatedAtColumnName)),
		UpdatedAt: yago.NewScalarField(mapper.Table().C(PersonUpdatedAtColumnName)),
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
		m[PersonIDColumnName] = s.ID
	}
	m[PersonNameColumnName] = s.Name
	m[PersonEmailColumnName] = s.Email
	m[PersonCreatedAtColumnName] = s.CreatedAt
	m[PersonUpdatedAtColumnName] = s.UpdatedAt
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personTable.C(PersonIDColumnName),
		personTable.C(PersonNameColumnName),
		personTable.C(PersonEmailColumnName),
		personTable.C(PersonCreatedAtColumnName),
		personTable.C(PersonUpdatedAtColumnName),
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

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PersonMapper) AutoIncrementPKey() bool {
	return true
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PersonMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	s := instance.(*Person)
	s.ID = value
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return personTable.C(PersonIDColumnName).Eq(instance.(*Person).ID)
}


const (
	PhoneNumberTableName = "phonenumber"
	PhoneNumberIDColumnName = "id"
	PhoneNumberPersonIDColumnName = "person_id"
	PhoneNumberNameColumnName = "name"
	PhoneNumberNumberColumnName = "number"
)

var phoneNumberTable = qb.Table(
	PhoneNumberTableName,
	qb.Column(PhoneNumberIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(PhoneNumberPersonIDColumnName, qb.BigInt()).NotNull(),
	qb.Column(PhoneNumberNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PhoneNumberNumberColumnName, qb.Varchar()).NotNull(),
	qb.ForeignKey().Ref(PhoneNumberPersonIDColumnName, PersonTableName, PersonIDColumnName),
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
		ID: yago.NewScalarField(mapper.Table().C(PhoneNumberIDColumnName)),
		PersonID: yago.NewScalarField(mapper.Table().C(PhoneNumberPersonIDColumnName)),
		Name: yago.NewScalarField(mapper.Table().C(PhoneNumberNameColumnName)),
		Number: yago.NewScalarField(mapper.Table().C(PhoneNumberNumberColumnName)),
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
		m[PhoneNumberIDColumnName] = s.ID
	}
	m[PhoneNumberPersonIDColumnName] = s.PersonID
	m[PhoneNumberNameColumnName] = s.Name
	m[PhoneNumberNumberColumnName] = s.Number
	return m
}

// FieldList returns the list of fields for a select
func (mapper PhoneNumberMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		phoneNumberTable.C(PhoneNumberIDColumnName),
		phoneNumberTable.C(PhoneNumberPersonIDColumnName),
		phoneNumberTable.C(PhoneNumberNameColumnName),
		phoneNumberTable.C(PhoneNumberNumberColumnName),
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

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PhoneNumberMapper) AutoIncrementPKey() bool {
	return true
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PhoneNumberMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	s := instance.(*PhoneNumber)
	s.ID = value
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PhoneNumberMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return phoneNumberTable.C(PhoneNumberIDColumnName).Eq(instance.(*PhoneNumber).ID)
}
