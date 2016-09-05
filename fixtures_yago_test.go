package yago_test

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
	"github.com/m4rw3r/uuid"
)

const (
	// SimpleStructTableName is the SimpleStruct associated table name
	SimpleStructTableName = "simplestruct"
	// SimpleStructIDColumnName is the ID field associated column name
	SimpleStructIDColumnName = "id"
	// SimpleStructNameColumnName is the Name field associated column name
	SimpleStructNameColumnName = "name"
)

var simpleStructTable = qb.Table(
	SimpleStructTableName,
	qb.Column(SimpleStructIDColumnName, qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column(SimpleStructNameColumnName, qb.Varchar()).NotNull(),
	qb.UniqueKey(
		SimpleStructNameColumnName,
	),
)

var simpleStructType = reflect.TypeOf(SimpleStruct{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (SimpleStruct) StructType() reflect.Type {
	return simpleStructType
}

// SimpleStructModel
type SimpleStructModel struct {
	mapper *SimpleStructMapper
	ID yago.ScalarField
	Name yago.ScalarField
}

func NewSimpleStructModel(meta *yago.Metadata) SimpleStructModel {
	mapper := NewSimpleStructMapper()
	meta.AddMapper(mapper)
	return SimpleStructModel {
		mapper: mapper,
		ID: yago.NewScalarField(mapper.Table().C(SimpleStructIDColumnName)),
		Name: yago.NewScalarField(mapper.Table().C(SimpleStructNameColumnName)),
	}
}

func (m SimpleStructModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewSimpleStructMapper initialize a NewSimpleStructMapper
func NewSimpleStructMapper() *SimpleStructMapper {
	m := &SimpleStructMapper{}
	return m
}

// SimpleStructMapper is the SimpleStruct mapper
type SimpleStructMapper struct{}

// Name returns the mapper name
func (*SimpleStructMapper) Name() string {
	return "yago_test/SimpleStruct"
}

// Table returns the mapper table
func (*SimpleStructMapper) Table() *qb.TableElem {
	return &simpleStructTable
}

// StructType returns the reflect.Type of the mapped structure
func (SimpleStructMapper) StructType() reflect.Type {
	return simpleStructType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper SimpleStructMapper) SQLValues(instance yago.MappedStruct) map[string]interface{} {
	s, ok := instance.(*SimpleStruct)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	m := make(map[string]interface{})
	if s.ID != 0 {
		m[SimpleStructIDColumnName] = s.ID
	}
	m[SimpleStructNameColumnName] = s.Name
	return m
}

// FieldList returns the list of fields for a select
func (mapper SimpleStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		simpleStructTable.C(SimpleStructIDColumnName),
		simpleStructTable.C(SimpleStructNameColumnName),
	}
}

// Scan a struct
func (mapper SimpleStructMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*SimpleStruct)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.ID,
		&s.Name,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (SimpleStructMapper) AutoIncrementPKey() bool {
	return true
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (SimpleStructMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	s := instance.(*SimpleStruct)
	s.ID = value
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper SimpleStructMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return simpleStructTable.C(SimpleStructIDColumnName).Eq(instance.(*SimpleStruct).ID)
}

const (
	// PersonStructTableName is the PersonStruct associated table name
	PersonStructTableName = "personstruct"
	// PersonStructIDColumnName is the ID field associated column name
	PersonStructIDColumnName = "id"
	// PersonStructFirstNameColumnName is the FirstName field associated column name
	PersonStructFirstNameColumnName = "first_name"
	// PersonStructLastNameColumnName is the LastName field associated column name
	PersonStructLastNameColumnName = "last_name"
)

var personStructTable = qb.Table(
	PersonStructTableName,
	qb.Column(PersonStructIDColumnName, qb.UUID()).PrimaryKey().NotNull(),
	qb.Column(PersonStructFirstNameColumnName, qb.Varchar()).NotNull(),
	qb.Column(PersonStructLastNameColumnName, qb.Varchar()).NotNull(),
)

var personStructType = reflect.TypeOf(PersonStruct{})

// StructType returns the reflect.Type of the struct
// It is used for indexing mappers (and only that I guess, so
// it could be replaced with a unique identifier).
func (PersonStruct) StructType() reflect.Type {
	return personStructType
}

// PersonStructModel
type PersonStructModel struct {
	mapper *PersonStructMapper
	ID yago.ScalarField
	FirstName yago.ScalarField
	LastName yago.ScalarField
}

func NewPersonStructModel(meta *yago.Metadata) PersonStructModel {
	mapper := NewPersonStructMapper()
	meta.AddMapper(mapper)
	return PersonStructModel {
		mapper: mapper,
		ID: yago.NewScalarField(mapper.Table().C(PersonStructIDColumnName)),
		FirstName: yago.NewScalarField(mapper.Table().C(PersonStructFirstNameColumnName)),
		LastName: yago.NewScalarField(mapper.Table().C(PersonStructLastNameColumnName)),
	}
}

func (m PersonStructModel) GetMapper() yago.Mapper {
	return m.mapper
}

// NewPersonStructMapper initialize a NewPersonStructMapper
func NewPersonStructMapper() *PersonStructMapper {
	m := &PersonStructMapper{}
	return m
}

// PersonStructMapper is the PersonStruct mapper
type PersonStructMapper struct{}

// Name returns the mapper name
func (*PersonStructMapper) Name() string {
	return "yago_test/PersonStruct"
}

// Table returns the mapper table
func (*PersonStructMapper) Table() *qb.TableElem {
	return &personStructTable
}

// StructType returns the reflect.Type of the mapped structure
func (PersonStructMapper) StructType() reflect.Type {
	return personStructType
}

// SQLValues returns values as a map
// The primary key is included only if having non-default values
func (mapper PersonStructMapper) SQLValues(instance yago.MappedStruct) map[string]interface{} {
	s, ok := instance.(*PersonStruct)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	m := make(map[string]interface{})
	if s.ID != (uuid.UUID{}) {
		m[PersonStructIDColumnName] = s.ID
	}
	m[PersonStructFirstNameColumnName] = s.FirstName
	m[PersonStructLastNameColumnName] = s.LastName
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personStructTable.C(PersonStructIDColumnName),
		personStructTable.C(PersonStructFirstNameColumnName),
		personStructTable.C(PersonStructLastNameColumnName),
	}
}

// Scan a struct
func (mapper PersonStructMapper) Scan(rows *sql.Rows, instance yago.MappedStruct) error {
	s, ok := instance.(*PersonStruct)
	if !ok {
		panic("Wrong struct type passed to the mapper")
	}
	return rows.Scan(
		&s.ID,
		&s.FirstName,
		&s.LastName,
	)
}

// AutoIncrementPKey return true if a column of the pkey is autoincremented
func (PersonStructMapper) AutoIncrementPKey() bool {
	return false
}

// LoadAutoIncrementPKeyValue set the pkey autoincremented column value
func (PersonStructMapper) LoadAutoIncrementPKeyValue(instance yago.MappedStruct, value int64) {
	panic("PersonStruct has no auto increment column in its pkey")
}

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonStructMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return personStructTable.C(PersonStructIDColumnName).Eq(instance.(*PersonStruct).ID)
}
