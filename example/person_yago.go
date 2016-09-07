package main

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
)

const (
	// PersonName is the Name field name
	PersonName = "Name"
	// PersonNameColumnName is the Name field associated column name
	PersonNameColumnName = "name"
	// PersonEmail is the Email field name
	PersonEmail = "Email"
	// PersonEmailColumnName is the Email field associated column name
	PersonEmailColumnName = "email_address"
)

const (
	// PersonTableName is the Person associated table name
	PersonTableName = "person"
)

var personTable = qb.Table(
	PersonTableName,
	qb.Column(PersonNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PersonEmailColumnName, qb.Varchar()).Null(),
	qb.Column(BaseIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(BaseCreatedAtColumnName, qb.Timestamp()).NotNull(),
	qb.Column(BaseUpdatedAtColumnName, qb.Timestamp()).Null(),
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

// PersonModel provides direct access to helpers for Person
// queries
type PersonModel struct {
	mapper    *PersonMapper
	Name      yago.ScalarField
	Email     yago.ScalarField
	ID        yago.ScalarField
	CreatedAt yago.ScalarField
	UpdatedAt yago.ScalarField
}

// NewPersonModel returns a new PersonModel
func NewPersonModel(meta *yago.Metadata) PersonModel {
	mapper := NewPersonMapper()
	meta.AddMapper(mapper)
	return PersonModel{
		mapper:    mapper,
		Name:      yago.NewScalarField(mapper.Table().C(PersonNameColumnName)),
		Email:     yago.NewScalarField(mapper.Table().C(PersonEmailColumnName)),
		ID:        yago.NewScalarField(mapper.Table().C(BaseIDColumnName)),
		CreatedAt: yago.NewScalarField(mapper.Table().C(BaseCreatedAtColumnName)),
		UpdatedAt: yago.NewScalarField(mapper.Table().C(BaseUpdatedAtColumnName)),
	}
}

// GetMapper returns the associated PersonMapper instance
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

// GetMapper returns itself
func (mapper *PersonMapper) GetMapper() yago.Mapper {
	return mapper
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

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper PersonMapper) SQLValues(instance yago.MappedStruct, fields ...string) map[string]interface{} {
	s, ok := instance.(*Person)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	allValues := len(fields) == 0
	m := make(map[string]interface{})
	if s.ID != 0 {
		m[BaseIDColumnName] = s.ID
	}
	if allValues || yago.StringListContains(fields, PersonName) {
		m[PersonNameColumnName] = s.Name
	}
	if allValues || yago.StringListContains(fields, PersonEmail) {
		m[PersonEmailColumnName] = s.Email
	}
	if allValues || yago.StringListContains(fields, BaseCreatedAt) {
		m[BaseCreatedAtColumnName] = s.CreatedAt
	}
	if allValues || yago.StringListContains(fields, BaseUpdatedAt) {
		m[BaseUpdatedAtColumnName] = s.UpdatedAt
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personTable.C(PersonNameColumnName),
		personTable.C(PersonEmailColumnName),
		personTable.C(BaseIDColumnName),
		personTable.C(BaseCreatedAtColumnName),
		personTable.C(BaseUpdatedAtColumnName),
	}
}

// Scan a struct
func (mapper PersonMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*Person)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.Name,
		&s.Email,
		&s.ID,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PersonMapper) AutoIncrementPKey() bool {
	return false
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PersonMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	panic("Person has no auto increment column in its pkey")
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return personTable.C(BaseIDColumnName).Eq(instance.(*Person).ID)
}

const (
	// PhoneNumberPersonID is the PersonID field name
	PhoneNumberPersonID = "PersonID"
	// PhoneNumberPersonIDColumnName is the PersonID field associated column name
	PhoneNumberPersonIDColumnName = "person_id"
	// PhoneNumberName is the Name field name
	PhoneNumberName = "Name"
	// PhoneNumberNameColumnName is the Name field associated column name
	PhoneNumberNameColumnName = "name"
	// PhoneNumberNumber is the Number field name
	PhoneNumberNumber = "Number"
	// PhoneNumberNumberColumnName is the Number field associated column name
	PhoneNumberNumberColumnName = "number"
)

const (
	// PhoneNumberTableName is the PhoneNumber associated table name
	PhoneNumberTableName = "phonenumber"
)

var phoneNumberTable = qb.Table(
	PhoneNumberTableName,
	qb.Column(PhoneNumberPersonIDColumnName, qb.BigInt()).NotNull(),
	qb.Column(PhoneNumberNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PhoneNumberNumberColumnName, qb.Varchar()).NotNull(),
	qb.Column(BaseIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(BaseCreatedAtColumnName, qb.Timestamp()).NotNull(),
	qb.Column(BaseUpdatedAtColumnName, qb.Timestamp()).Null(),
	qb.ForeignKey(PhoneNumberPersonIDColumnName).References(PersonTableName, BaseIDColumnName).OnUpdate("CASCADE").OnDelete("CASCADE"),
)

var phoneNumberType = reflect.TypeOf(PhoneNumber{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (PhoneNumber) StructType() reflect.Type {
	return phoneNumberType
}

// PhoneNumberModel provides direct access to helpers for PhoneNumber
// queries
type PhoneNumberModel struct {
	mapper    *PhoneNumberMapper
	PersonID  yago.ScalarField
	Name      yago.ScalarField
	Number    yago.ScalarField
	ID        yago.ScalarField
	CreatedAt yago.ScalarField
	UpdatedAt yago.ScalarField
}

// NewPhoneNumberModel returns a new PhoneNumberModel
func NewPhoneNumberModel(meta *yago.Metadata) PhoneNumberModel {
	mapper := NewPhoneNumberMapper()
	meta.AddMapper(mapper)
	return PhoneNumberModel{
		mapper:    mapper,
		PersonID:  yago.NewScalarField(mapper.Table().C(PhoneNumberPersonIDColumnName)),
		Name:      yago.NewScalarField(mapper.Table().C(PhoneNumberNameColumnName)),
		Number:    yago.NewScalarField(mapper.Table().C(PhoneNumberNumberColumnName)),
		ID:        yago.NewScalarField(mapper.Table().C(BaseIDColumnName)),
		CreatedAt: yago.NewScalarField(mapper.Table().C(BaseCreatedAtColumnName)),
		UpdatedAt: yago.NewScalarField(mapper.Table().C(BaseUpdatedAtColumnName)),
	}
}

// GetMapper returns the associated PhoneNumberMapper instance
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

// GetMapper returns itself
func (mapper *PhoneNumberMapper) GetMapper() yago.Mapper {
	return mapper
}

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
func (mapper PhoneNumberMapper) SQLValues(instance yago.MappedStruct, fields ...string) map[string]interface{} {
	s, ok := instance.(*PhoneNumber)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	allValues := len(fields) == 0
	m := make(map[string]interface{})
	if s.ID != 0 {
		m[BaseIDColumnName] = s.ID
	}
	if allValues || yago.StringListContains(fields, PhoneNumberPersonID) {
		m[PhoneNumberPersonIDColumnName] = s.PersonID
	}
	if allValues || yago.StringListContains(fields, PhoneNumberName) {
		m[PhoneNumberNameColumnName] = s.Name
	}
	if allValues || yago.StringListContains(fields, PhoneNumberNumber) {
		m[PhoneNumberNumberColumnName] = s.Number
	}
	if allValues || yago.StringListContains(fields, BaseCreatedAt) {
		m[BaseCreatedAtColumnName] = s.CreatedAt
	}
	if allValues || yago.StringListContains(fields, BaseUpdatedAt) {
		m[BaseUpdatedAtColumnName] = s.UpdatedAt
	}
	return m
}

// FieldList returns the list of fields for a select
func (mapper PhoneNumberMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		phoneNumberTable.C(PhoneNumberPersonIDColumnName),
		phoneNumberTable.C(PhoneNumberNameColumnName),
		phoneNumberTable.C(PhoneNumberNumberColumnName),
		phoneNumberTable.C(BaseIDColumnName),
		phoneNumberTable.C(BaseCreatedAtColumnName),
		phoneNumberTable.C(BaseUpdatedAtColumnName),
	}
}

// Scan a struct
func (mapper PhoneNumberMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*PhoneNumber)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.PersonID,
		&s.Name,
		&s.Number,
		&s.ID,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PhoneNumberMapper) AutoIncrementPKey() bool {
	return false
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PhoneNumberMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	panic("PhoneNumber has no auto increment column in its pkey")
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PhoneNumberMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return phoneNumberTable.C(BaseIDColumnName).Eq(instance.(*PhoneNumber).ID)
}
