package yago_test

// generated with yago. Better NOT to edit!

import (
	"database/sql"
	"reflect"

	"github.com/aacanakin/qb"

	"github.com/orus-io/yago"
	"github.com/m4rw3r/uuid"
)



var simpleStructTable = qb.Table(
	"simplestruct",
	qb.Column("id", qb.BigInt()).PrimaryKey().AutoIncrement().NotNull(),
	qb.Column("name", qb.Varchar()).NotNull(),
	qb.UniqueKey(
		"name",
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
		ID: yago.NewScalarField(mapper.Table().C("id")),
		Name: yago.NewScalarField(mapper.Table().C("name")),
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
		m["id"] = s.ID
	}
	m["name"] = s.Name
	return m
}

// FieldList returns the list of fields for a select
func (mapper SimpleStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		simpleStructTable.C("id"),
		simpleStructTable.C("name"),
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

// PKeyClause returns a clause that matches the instance primary key
func (mapper SimpleStructMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return simpleStructTable.C("id").Eq(instance.(*SimpleStruct).ID)
}


var personStructTable = qb.Table(
	"personstruct",
	qb.Column("id", qb.UUID()).PrimaryKey().NotNull(),
	qb.Column("first_name", qb.Varchar()).NotNull(),
	qb.Column("last_name", qb.Varchar()).NotNull(),
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
		ID: yago.NewScalarField(mapper.Table().C("id")),
		FirstName: yago.NewScalarField(mapper.Table().C("first_name")),
		LastName: yago.NewScalarField(mapper.Table().C("last_name")),
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
		m["id"] = s.ID
	}
	m["first_name"] = s.FirstName
	m["last_name"] = s.LastName
	return m
}

// FieldList returns the list of fields for a select
func (mapper PersonStructMapper) FieldList() []qb.Clause {
	return []qb.Clause{
		personStructTable.C("id"),
		personStructTable.C("first_name"),
		personStructTable.C("last_name"),
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

// PKeyClause returns a clause that matches the instance primary key
func (mapper PersonStructMapper) PKeyClause(instance yago.MappedStruct) qb.Clause {
	return personStructTable.C("id").Eq(instance.(*PersonStruct).ID)
}
